package main

import (
    "io/ioutil"
    "log"
)

func create() {
    sc, err := NewSpannerClient() 
    if err != nil {
		log.Fatalf("Couldn't connect to Google Cloud Spanner: %v", err)
	}
    defer sc.disconnect()

    file, err := ioutil.ReadFile(*schemaFile)
	if err != nil {
		log.Fatal(err)
	}

    err = sc.createDB(string(file))
	if err != nil {
		log.Fatal(err)
	}
}
