package main

import (
    "log"
)

func generate() {
    // TODO: Implement.
    gcs, err := NewGCSClient()
    if err != nil {
        log.Fatalf("Couldn't connect to Google Cloud Storage: %v", err)
    }

    gcs.list(*bucketName, "")
}
