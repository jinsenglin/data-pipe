package main

import (
    "context"
    "log"
    "sync"
    "time"

    "cloud.google.com/go/storage"
    "github.com/gocarina/gocsv"
)

// GCSclient is an authenticated Cloud Storage client.
// Golang technique here is embedding.
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

// write for passed filePath
func (client *GCSclient) write (bucketName, filePath string, fileContent []byte) (error) {
    ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
    defer cancel()

    writer := client.Bucket(bucketName).Object(filePath).NewWriter(ctx)
    defer writer.Close()

    _, err := writer.Write(fileContent)
    if err != nil {
		return err
	}

    return nil
}

// marshal passed objects then write for passed filePath
func (client *GCSclient) writeCSV(bucketName, filePath string, objects interface{}, wg *sync.WaitGroup) error {
    defer wg.Done()

    debugger.Printf("Writing csv %s", filePath)

    csv, err := gocsv.MarshalString(objects)
    if err != nil {
		return err
    }

    err = client.write(bucketName, filePath, []byte(csv))
    if err != nil {
        log.Printf("Couldn't write csv %s: %v", filePath, err)
		return err
    }

    debugger.Printf("Done writing csv %s", filePath)

    return nil
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
