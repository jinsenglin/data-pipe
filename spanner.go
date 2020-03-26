package main

import (
    "context"
    "fmt"
    "strings"
    "time"

    "cloud.google.com/go/spanner"
    admindbapi "cloud.google.com/go/spanner/admin/database/apiv1"
    admindbpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

// Spannerclient is an authenticated Cloud Spanner client.
// Golang technique: type embedding.
type SpannerClient struct {
    *spanner.Client
    admin           *admindbapi.DatabaseAdminClient
    instancePath    string
    databasePath    string
}

// NewSpannerClient creates new authenticated Cloud Spanner client.
// The client will use your default application credentials.
func NewSpannerClient() (*SpannerClient, error) {
    instancePath := fmt.Sprintf("projects/%s/instances/%s", *project, *instance)
    databasePath := fmt.Sprintf("%s/databases/%s", instancePath, *database)

    debugger.Printf("Connecting Spanner client to %s", databasePath)

    ctx := context.Background()

    // TODO: client, err := spanner.NewClient(ctx, databasePath, spanner.ClientConfig{NumChannels: 20})
    client, err := spanner.NewClient(ctx, databasePath)
	if err != nil {
		return nil, err
	}

    admin, err := admindbapi.NewDatabaseAdminClient(ctx)
	if err != nil {
		return nil, err
	}

	return &SpannerClient{client, admin, instancePath, databasePath}, nil
}

func (client *SpannerClient) createDB (schema string) error {
    statementsRaw := strings.Split(schema, ";")

	statements := make([]string, 0, len(statementsRaw))
	for i := 0; i < len(statementsRaw); i++ {
		statement := strings.TrimSpace(statementsRaw[i])
		if len(statement) > 0 && !strings.HasPrefix(statement, "--") {
			statements = append(statements, statement)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	op, err := client.admin.CreateDatabase(ctx, &admindbpb.CreateDatabaseRequest{
		Parent:          client.instancePath,
		CreateStatement: "CREATE DATABASE `" + *database + "`",
		ExtraStatements: statements,
	})
	if err != nil {
		return err
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return err
	}

    return nil
}

func (client *SpannerClient) cleanup () {
    debugger.Println("Closing Google Cloud Spanner connections...")
	client.Close()
	client.admin.Close()
	debugger.Println("Finished closing Google Cloud Spanner connections.")
}

func (client *SpannerClient) newMutation (table string, s interface{}) (*spanner.Mutation, error) {
    mutation, err := spanner.InsertOrUpdateStruct(table, s)
    if err != nil {
        return nil, err
    }

    return mutation, nil
}

func (client *SpannerClient) write (mutations []*spanner.Mutation) (error) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if _, err := client.Apply(ctx, mutations); err != nil {
        return err
    }

    return nil
}


