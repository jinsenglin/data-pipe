package main

import (
    "context"
    "io/ioutil"
    "log"
    "sort"
    "sync"
    "time"

    "cloud.google.com/go/storage"
    "github.com/gocarina/gocsv"
    "google.golang.org/api/iterator"
)

// GCSclient is an authenticated Cloud Storage client.
// Golang technique here is embedding.
type GCSclient struct {
    *storage.Client
}

// list for passed bucketName filtered by passed filePrefix
func (client *GCSclient) list (bucketName string, filePrefix string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var files []string

	it := client.Bucket(bucketName).Objects(ctx, &storage.Query{Prefix: filePrefix})

	for {
		objAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		files = append(files, objAttrs.Name)
	}

	sort.Strings(files)

    return files, nil
}

// read for passed filePath
func (client *GCSclient) read (bucketName, filePath string) ([]byte, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	reader, err := client.Bucket(bucketName).Object(filePath).NewReader(ctx)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	file, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

    return file, nil
}

// read for passed filePath then unmarshal the file content
func (client *GCSclient) readCSV (bucketName, filePath string, callback func(interface{}) error) (error) {
    // TODO: Implement.

    file, err := client.read(bucketName, filePath)

    if err != nil {
        return err
    }

    gocsv.UnmarshalBytesToCallback(file, callback)

/*
                    gocsv.UnmarshalBytesToCallback(j.files[j.mapping["accounts"]], func(a *dbAccount) error {
						id := uuid.New()
						a.AccountID = string(encodeHexUUID(&id))
						j.lookup["accounts"][a.OldID] = &id
						return appendMutation(chs[0], "Account", a.AccountID, a)
					})
*/

    return nil
}

// write for passed filePath
func (client *GCSclient) write (bucketName, filePath string, fileContent []byte) (error) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    writer := client.Bucket(bucketName).Object(filePath).NewWriter(ctx)

    if _, err := writer.Write(fileContent); err != nil {
		return err
	}

    if err := writer.Close(); err != nil {
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

    if err := client.write(bucketName, filePath, []byte(csv)); err != nil {
        log.Printf("Couldn't write csv %s: %v", filePath, err)
		return err
    }

    debugger.Printf("Done writing csv %s", filePath)

    return nil
}

// NewGCSClient creates new authenticated Cloud Storage client.
// The client will use your default application credentials.
func NewGCSClient() (*GCSclient, error) {
    ctx := context.Background()

    client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &GCSclient{client}, nil
}
