package main

import (
    "log"
)

func reset() {
    sc, err := NewSpannerClient()
    if err != nil {
		log.Fatalf("Couldn't connect to Google Cloud Spanner: %v", err)
	}

    debugger.Printf("Deleting all rows in %s table...", "Account")
    if err = sc.clearTable("Account"); err != nil {
        log.Fatalf("Couldn't delete rows from %s table: %v", "Account", err)
	}

    debugger.Printf("Deleting all rows in %s table...", "Singer")
    if err = sc.clearTable("Singer"); err != nil {
        log.Fatalf("Couldn't delete rows from %s table: %v", "Singer", err)
	}
}
