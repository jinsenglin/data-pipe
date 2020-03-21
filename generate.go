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

    wgGenerate := &sync.WaitGroup{}

    wgGenerate.Add(1)
    go generateAccounts(gcs, wgGenerate)
//    wgGenerate.Add(1)
//    go generateSigners(gcs, wgGenerate)
//    wgGenerate.Add(1)
//    go generateAlbums(gcs, wgGenerate)
//    wgGenerate.Add(1)
//    go generateSongs(gcs, wgGenerate)

    debugger.Println("Waiting for write to GCS to finish....")
    wgGenerate.Wait()
    debugger.Println("Done generating")
}

func generateAccounts(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v accounts...", *numAccounts)

    wgGenerateAccounts := &sync.WaitGroup{}

    var accounts []*Account
    for i := 0; i < *numAccounts; i++ {
        accounts = append(accounts, NewAccount(int64(i)))

        if len(accounts) % recordsPerFile == 0 {
            fileName := fmt.Sprintf("%v-%04d.csv", "account", i/recordsPerFile)

            wgGenerateAccounts.Add(1)
            go gcs.writeCSV(*bucketName, fileName, accounts, wgGenerateAccounts)

            accounts = []*Account{}
        }
    }

    fileName := fmt.Sprintf("%v-%04d.csv", "account", *numAccounts/recordsPerFile)

    wgGenerateAccounts.Add(1)
    go gcs.writeCSV(*bucketName, fileName, accounts, wgGenerateAccounts)

    debugger.Println("Waiting for write accounts to GCS to finish....")
    wgGenerateAccounts.Wait()

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
