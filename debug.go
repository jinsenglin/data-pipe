package main

import (
    "log"
)

type debugging bool

func (d debugging) Printf(format string, args ...interface{}) {
	if d {
		log.Printf(format, args...)
	}
}

func (d debugging) Print(args ...interface{}) {
	if d {
		log.Print(args...)
	}
}

func (d debugging) Println(args ...interface{}) {
	if d {
		log.Println(args...)
	}
}

func (d debugging) DumpVariables() {
	if d {
		log.Println("Dump Variables")
		log.Printf("%-10s = %v", "version", version)
		log.Printf("%-10s = %v", "project", *project)
		log.Printf("%-10s = %v", "bucketName", *bucketName)
	}
}


