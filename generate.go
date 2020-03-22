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

//    wgGenerate.Add(1)
//    go generateAccounts(gcs, wgGenerate)
    wgGenerate.Add(1)
    go generateSingers(gcs, wgGenerate)
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

    if len(accounts) > 0 {
        fileName := fmt.Sprintf("%v-%04d.csv", "account", *numAccounts/recordsPerFile)

        wgGenerateAccounts.Add(1)
        go gcs.writeCSV(*bucketName, fileName, accounts, wgGenerateAccounts)
    }

    debugger.Println("Waiting for write accounts to GCS to finish....")
    wgGenerateAccounts.Wait()

    debugger.Printf("Done generating %v accounts...", *numAccounts)
}

func generateSingers(gcs *GCSclient, wg *sync.WaitGroup) {
	defer wg.Done()

    debugger.Printf("Generating %v singers...", *numSingers)

    wgGenerateSingers := &sync.WaitGroup{}

    chSingers := make(chan *Singer, *numSingers)

    wgGenerateSingers.Add(1)
    go generateSingersConsumer(gcs, chSingers, wgGenerateSingers)

    batchSingers := 0
    for batchSingers = *numSingers; batchSingers > recordsPerFile; batchSingers -= recordsPerFile {
        wgGenerateSingers.Add(1)
        go generateSingersProducer(recordsPerFile, chSingers, wgGenerateSingers)
    }
    if batchSingers > 0 {
        go generateSingersProducer(batchSingers, chSingers, wgGenerateSingers)
    }

    debugger.Println("Waiting for write signers to GCS to finish....")
    wgGenerateSingers.Wait()
    close(chSingers)

    debugger.Printf("Done generating %v singers...", *numSingers)
}

func generateSingersProducer(batchSize int, out chan<- *Singer, wg *sync.WaitGroup) {
    defer wg.Done()

    debugger.Printf("Producing %v singers...", batchSize)

    for i := 0; i < batchSize; i++ {
        out <- NewSinger("") // TODO: Implement uuid v4 id
    }

    debugger.Printf("Done producing %v singers...", batchSize)
}
func generateSingersConsumer(gcs *GCSclient, in <-chan *Singer, wg *sync.WaitGroup) {
    defer wg.Done()

    debugger.Printf("Consuming singers...")

    for singer := range in {
        // TODO: Implement write signer
        if singer != nil {
        }
    }

    debugger.Printf("Done consuming singers...")
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
