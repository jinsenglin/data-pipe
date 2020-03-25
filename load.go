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

    // TODO: Implement.

    ch := make(chan *Account, *numAccounts)
    go func() {
        // TODO: Implement.
        account := <- ch
        debugger.PrintNothing(account)
    }()
    for _, file := range files {
        debugger.Printf("Loading file %s...", file)
        gcs.readCSV(*bucketName, file, func(account *Account) {
            // TODO: Implement.
            ch <- account
        })
    }
}
