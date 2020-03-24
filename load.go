package main

import (
    "log"
)

func load() {
    gcs, err := NewGCSClient()
    if err != nil {
        log.Fatalf("Couldn't connect to Google Cloud Storage: %v", err)
    }

    files, err := gcs.list(*bucketName, "account-")
    if err != nil {
        log.Fatalf("Couldn't list files of prefix account- %v", err)
    }

    for _, file := range files {
        debugger.Printf("Loading file %s...", file)
        gcs.readCSV(*bucketName, file, func(account *Account) {
            // TODO: Implement.
        })
    }
}
