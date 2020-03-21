package main

import (
    "log"

    "github.com/namsral/flag"
)

var (
    version string // set by linker -X

    project = flag.String("project", "", "Your cloud project ID.")
    bucketName = flag.String("bucket", "", "The name of the bucket within your project.")

    debugger debugging
)

func main() {
    // init var
    flag.Parse()

    // init var
    debugger = debugging(true)

    debugger.DumpVariables()

    gcs, err := NewGCSClient()
    if err != nil {
        log.Fatalf("Couldn't connect to Google Cloud Storage: %v", err)
    }

    gcs.list(*bucketName, "")
}
