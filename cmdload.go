package main

import (
    "log"
    "sync"

    "cloud.google.com/go/spanner"
)

func load() {
    sc, err := NewSpannerClient()
    if err != nil {
        log.Fatalf("Couldn't connect to Google Cloud Spanner: %v", err)
    }

    gcs, err := NewGCSClient()
    if err != nil {
        log.Fatalf("Couldn't connect to Google Cloud Storage: %v", err)
    }

    loadAccount(sc, gcs)
}

func loadAccount(sc *SpannerClient, gcs *GCSclient) {
    files, err := gcs.list(*bucket, "account-")
    if err != nil {
        log.Fatalf("Couldn't list files of prefix account- %v", err)
    }

    chAccount := make(chan *Account, *numAccounts)
    chMutation := make(chan *spanner.Mutation, *numAccounts)
    chMutations := make(chan []*spanner.Mutation, *numAccounts/100)

    wgLoaders := &sync.WaitGroup{}
    for i := 0; i < *numLoaders; i++ {
        debugger.Printf("Starting loader %d", i)

        wgLoaders.Add(1)
        go func(in <- chan []*spanner.Mutation) {
            defer wgLoaders.Done()
            i := 1
            for mutations := range in {
                if err := sc.write(mutations); err != nil {
                    // TODO: Implement retry.
                    log.Fatalf("%v", err)
                }
                debugger.Printf("Applied batch %d", i)
                i++
            }
        }(chMutations)
    }

    wgBatchMakers := &sync.WaitGroup{}
    // TODO: for i := 0; i < *numAccounts/recordsPerFile; i++ {
    for i := 0; i < 1; i++ {
        debugger.Printf("Starting batch maker %d", i)

        wgBatchMakers.Add(1)
        go func(in <- chan *spanner.Mutation, out chan <- []*spanner.Mutation) {
            defer wgBatchMakers.Done()
            for mutation := range in {
                // TODO: Implement batch size.
                out <- []*spanner.Mutation{mutation}
            }
        }(chMutation, chMutations)
    }

    wgMutationMakers := &sync.WaitGroup{}
    // TODO: for i := 0; i < *numAccounts/recordsPerFile; i++ {
    for i := 0; i < 1; i++ {
        debugger.Printf("Starting mutation maker %d", i)

        wgMutationMakers.Add(1)
        go func(in <- chan *Account, out chan <- *spanner.Mutation) {
            defer wgMutationMakers.Done()
            for account := range in {
                mutation, _ := sc.newMutation("Account", account)
                // TODO: Implement error handling.
                out <- mutation
            }
        }(chAccount, chMutation)
    }

    wgReaders := &sync.WaitGroup{}
    for i, file := range files {
        debugger.Printf("Starting reader %d for file %s...", i, file)

        wgReaders.Add(1)
        go func(out chan <- *Account) {
            defer wgReaders.Done()
            gcs.readCSV(*bucket, file, func(account *Account) {out <- account})
        }(chAccount)
    }

    debugger.Println("loadAccount is waiting for readers...")
    wgReaders.Wait()
    close(chAccount)

    debugger.Println("loadAccount is waiting for mutation makers...")
    wgMutationMakers.Wait()
    close(chMutation)

    debugger.Println("loadAccount is waiting for batch makers...")
    wgBatchMakers.Wait()
    close(chMutations)

    debugger.Println("loadAccount is waiting for loaders...")
    wgLoaders.Wait()

    debugger.Println("loadAccount is done.")
}
