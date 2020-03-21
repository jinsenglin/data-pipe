package main

import (
    "context"
    "time"

    "cloud.google.com/go/storage"
)

// GCSclient is an authenticated Cloud Storage client
type GCSclient struct {
    *storage.Client
}

// list for passed bucketName filtered by passed filePrefix
func (client *GCSclient) list (bucketName string, filePrefix string) ([]string, error) {
    var files []string

    // TODO: Implement.

    return files, nil
}

// read for passed filePath
func (client *GCSclient) read (bucketName, filePath string) ([]byte, error) {
    var file []byte

    // TODO: Implement

    return file, nil
}

// NewGCSService creates new authenticated Cloud Storage client.
// The client will use your default application credentials.
func NewGCSClient() (*GCSclient, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &GCSclient{client}, nil
}
