package main

import (
    "fmt"
    "log"
    "sync"

    "github.com/google/uuid"
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

    wgGenerateSingersP := &sync.WaitGroup{}
    wgGenerateSingersC := &sync.WaitGroup{}

    chSingers := make(chan *Singer, *numSingers)

    wgGenerateSingersC.Add(1)
    go generateSingersConsumer(gcs, chSingers, wgGenerateSingersC)

    batchSingers := 0
    for batchSingers = *numSingers; batchSingers > recordsPerFile; batchSingers -= recordsPerFile {
        wgGenerateSingersP.Add(1)
        go generateSingersProducer(recordsPerFile, chSingers, wgGenerateSingersP)
    }
    if batchSingers > 0 {
        wgGenerateSingersP.Add(1)
        go generateSingersProducer(batchSingers, chSingers, wgGenerateSingersP)
    }

    debugger.Println("Waiting for write singers to GCS to finish....")
    wgGenerateSingersP.Wait()
    close(chSingers)
    wgGenerateSingersC.Wait()

    debugger.Printf("Done generating %v singers...", *numSingers)
}

func generateSingersProducer(batchSize int, out chan<- *Singer, wg *sync.WaitGroup) {
    defer wg.Done()

    debugger.Printf("Producing %v singers...", batchSize)

    for i := 0; i < batchSize; i++ {
        id := uuid.New().String()
        out <- NewSinger(id)
    }

    debugger.Printf("Done producing %v singers...", batchSize)
}
func generateSingersConsumer(gcs *GCSclient, in <-chan *Singer, wg *sync.WaitGroup) {
    defer wg.Done()

    debugger.Println("Consuming singers...")

    wgGenerateSingersConsumer := &sync.WaitGroup{}

    i := -1
    var singers []*Singer
    for singer := range in {
        i++
        singers = append(singers, singer)

        if len(singers) % recordsPerFile == 0 {
            fileName := fmt.Sprintf("%v-%04d.csv", "singer", i/recordsPerFile)

            wgGenerateSingersConsumer.Add(1)
            go gcs.writeCSV(*bucketName, fileName, singers, wgGenerateSingersConsumer)

            singers = []*Singer{}
        }
    }

    if len(singers) > 0 {
        fileName := fmt.Sprintf("%v-%04d.csv", "singer", *numAccounts/recordsPerFile)

        wgGenerateSingersConsumer.Add(1)
        go gcs.writeCSV(*bucketName, fileName, singers, wgGenerateSingersConsumer)
    }

    debugger.Println("Waiting for write singers to GCS to finish....")
    wgGenerateSingersConsumer.Wait()

    debugger.Println("Done consuming singers...")
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
