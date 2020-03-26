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

    chAccounts := make(chan *Account, *numAccounts)
    chMutations := make(chan *spanner.Mutation, *numAccounts)
    chBatchMutations := make(chan []*spanner.Mutation, *numAccounts/100)

    wgReaders := &sync.WaitGroup{}
    for i, file := range files {
        debugger.Printf("Starting reader %d for file %s...", i, file)

        wgReaders.Add(1)
        go func() {
            defer wgReaders.Done()
            gcs.readCSV(*bucket, file, func(account *Account) {chAccounts <- account})
        }()
    }

    wgMutationMakers := &sync.WaitGroup{}
    // TODO: for i := 0; i < *numAccounts/recordsPerFile; i++ {
    for i := 0; i < 1; i++ {
        debugger.Printf("Starting mutation maker %d", i)

        wgMutationMakers.Add(1)
        go func() {
            defer wgMutationMakers.Done()
            for account := range chAccounts {
                mutation, _ := sc.newMutation("Account", account)
                chMutations <- mutation
            }
        }()
    }

    wgBatchMakers := &sync.WaitGroup{}
    // TODO: for i := 0; i < *numAccounts/recordsPerFile; i++ {
    for i := 0; i < 1; i++ {
        debugger.Printf("Starting batch maker %d", i)

        wgBatchMakers.Add(1)
        go func() {
            defer wgBatchMakers.Done()
            for mutation := range chMutations {
                chBatchMutations <- []*spanner.Mutation{mutation}
            }
        }()
    }

    wgLoaders := &sync.WaitGroup{}
    for i := 0; i < *numLoaders; i++ {
        debugger.Printf("Starting loader %d", i)

        wgLoaders.Add(1)
        go func() {
            defer wgLoaders.Done()
            i := 1
            for batchMutation := range chBatchMutations {
                if err := sc.write(batchMutation); err != nil {
                    log.Fatalf("%v", err)
                }
                debugger.Printf("Applied batch %d", i)
                i++
            }
        }()
    }

    debugger.Println("loadAccount is waiting for readers...")
    wgReaders.Wait()
    close(chAccounts)

    debugger.Println("loadAccount is waiting for mutation makers...")
    wgMutationMakers.Wait()
    close(chMutations)

    debugger.Println("loadAccount is waiting for batch makers...")
    wgBatchMakers.Wait()
    close(chBatchMutations)

    debugger.Println("loadAccount is waiting for loaders...")
    wgLoaders.Wait()

    debugger.Println("loadAccount is done.")
}
