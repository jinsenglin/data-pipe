package main

import (
    "fmt"
    "context"

    "cloud.google.com/go/spanner"
)

// Spannerclient is an authenticated Cloud Spanner client.
// Golang technique here is embedding.
type Spannerclient struct {
    *spanner.Client
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
