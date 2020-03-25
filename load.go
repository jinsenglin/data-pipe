package main

import (
    "log"

    "cloud.google.com/go/spanner"
)

func load() {
    spannerClient, err := NewSpannerClient()
    if err != nil {
        log.Fatalf("Couldn't connect to Google Cloud Spanner: %v", err)
    }

    gcs, err := NewGCSClient()
    if err != nil {
        log.Fatalf("Couldn't connect to Google Cloud Storage: %v", err)
    }

    files, err := gcs.list(*bucketName, "account-")
    if err != nil {
        log.Fatalf("Couldn't list files of prefix account- %v", err)
    }

    // TODO: Implement.
    chAccounts := make(chan *Account, *numAccounts)
    chMutations := make(chan *spanner.Mutation, *numAccounts)
    chBatchMutations := make(chan []*spanner.Mutation, *numAccounts/100)

    for _, file := range files {
        debugger.Printf("Loading file %s...", file)
        gcs.readCSV(*bucketName, file, func(account *Account) {
            // TODO: Implement.
            chAccounts <- account
        })
    }

    for i := 0; i < *numAccounts/recordsPerFile; i++ {
        go func() {
            for account := range chAccounts {
                mutation, _ := spannerClient.newMutation("accounts", account)
                chMutations <- mutation
            }
        }()
    }

    for i := 0; i < *numAccounts/recordsPerFile; i++ {
        go func() {
            for mutation := range chMutations {
                chBatchMutations <- []*spanner.Mutation{mutation}
            }
        }()
    }

    for i := 0; i < *numLoaders; i++ {
        go func() {
            for batchMutation := range chBatchMutations {
                debugger.PrintNothing(batchMutation)
            }
        }()
    }

}
