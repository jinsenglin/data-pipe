package main

import (
    "fmt"
    "log"

    "github.com/namsral/flag"
)

var (
    version string // set by linker -X

    project = flag.String("project", "", "Your cloud project ID.")
    bucketName = flag.String("bucket", "", "The name of the bucket within your project.")
)

func main() {
    flag.Parse()

    fmt.Println(*project)
    fmt.Println(*bucketName)

    gcs, err := NewGCSClient()
    if err != nil {
        log.Fatalf("Couldn't connect to Google Cloud Storage: %v", err)
    }

    gcs.list(*bucketName, "")
}
