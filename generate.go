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
}
func generateSigners(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v signers...", 1000)

    // TODO: Implement.
}
func generateAlbums(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v albums...", 1000)

    // TODO: Implement.
}
func generateSongs(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v songs...", 1000)

    // TODO: Implement.
}
