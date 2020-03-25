package main

import (
    "context"
    "fmt"
    "time"

    "cloud.google.com/go/spanner"
)

// Spannerclient is an authenticated Cloud Spanner client.
// Golang technique here is embedding.
type Spannerclient struct {
    *spanner.Client
}

func (client *Spannerclient) newMutation (table string, s interface{}) (*spanner.Mutation, error) {
    mutation, err := spanner.InsertOrUpdateStruct(table, s)
    if err != nil {
        return nil, err
    }

    return mutation, nil
}

func (client *Spannerclient) write (mutations []*spanner.Mutation) (error) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if _, err := client.Apply(ctx, mutations); err != nil {
        return err
    }

    return nil
}

func (client *Spannerclient) cleanup () {
    // TODO: Implement.
}

func (client *Spannerclient) createDB (schema string) error {
    // TODO: Implement.
    return nil
}


// NewSpannerClient creates new authenticated Cloud Spanner client.
// The client will use your default application credentials.
func NewSpannerClient() (*Spannerclient, error) {
    ctx := context.Background()
    db := fmt.Sprintf("projects/%s/instances/%s/database/%s", *project, *instance, *database)

    client, err := spanner.NewClient(ctx, db)
	if err != nil {
		return nil, err
	}

	return &Spannerclient{client}, nil
}
