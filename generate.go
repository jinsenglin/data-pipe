package main

import (
    "fmt"
    "log"
    "sync"
)

func generate() {
    gcs, err := NewGCSClient()
    if err != nil {
        log.Fatalf("Couldn't connect to Google Cloud Storage: %v", err)
    }

    wg := &sync.WaitGroup{}

    wg.Add(1)
    go generateAccounts(gcs, wg)
    wg.Add(1)
    go generateSigners(gcs, wg)
    wg.Add(1)
    go generateAlbums(gcs, wg)
    wg.Add(1)
    go generateSongs(gcs, wg)

    debugger.Println("Waiting for write to GCS to finish....")
    wg.Wait()
}

func generateAccounts(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v accounts...", *numAccounts)

    var accounts []*Account
    for i := 0; i < *numAccounts; i++ {
        accounts = append(accounts, NewAccount(int64(i)))

        if len(accounts) % recordsPerFile == 0 {
            fileName := fmt.Sprintf("%v-%04d.csv", "account", i/recordsPerFile)
            err := gcs.writeCSV(*bucketName, fileName, accounts)
            if err != nil {
                log.Fatalf("Couldn't generate accounts: %v", err)
            }
        }

        accounts = []*Account{} 
    }
    fileName := fmt.Sprintf("%v-%04d.csv", "account", *numAccounts/recordsPerFile)
    err := gcs.writeCSV(*bucketName, fileName, accounts)
    if err != nil {
        log.Fatalf("Couldn't generate accounts: %v", err)
    }

    debugger.Printf("Done generating %v accounts...", *numAccounts)
}

func generateSigners(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v signers...", *numSigners)

    // TODO: Implement.

    debugger.Printf("Done generating %v signers...", *numSigners)
}

func generateAlbums(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v albums...", *numAlbums)

    // TODO: Implement.

    debugger.Printf("Done generating %v albums...", *numAlbums)
}

func generateSongs(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v songs...", *numSongs)

    // TODO: Implement.

    debugger.Printf("Done generating %v songs...", *numSongs)
}
