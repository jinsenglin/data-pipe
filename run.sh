#!/bin/bash

# By default, Go programs run with GOMAXPROCS set to the number of cores available
go run *.go -project gcp-expert-sandbox-jim -bucket gcp-expert-sandbox-jim-data-pipe -instance data-pipe -database data-pipe load
