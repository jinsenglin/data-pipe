package main

import (
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

    debugger.Printf("Generating %v accounts...", 1000)

    // TODO: Implement.
    for i := 0; i < *numAccounts; i++ {
        debugger.Printf("Generating account %v", i)
    }

    debugger.Printf("Done generating %v accounts...", 1000)
}
func generateSigners(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v signers...", 1000)

    // TODO: Implement.

    debugger.Printf("Done generating %v signers...", 1000)
}
func generateAlbums(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v albums...", 1000)

    // TODO: Implement.

    debugger.Printf("Done generating %v albums...", 1000)
}
func generateSongs(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v songs...", 1000)

    // TODO: Implement.

    debugger.Printf("Done generating %v songs...", 1000)
}
