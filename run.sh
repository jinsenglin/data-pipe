#!/bin/bash

# Setup
# gsutil mb -b on gs://gcp-expert-sandbox-jim-data-pipe
# gcloud spanner instances create data-pipe --config=regional-asia-east1 --description=DESCRIPTION --nodes=1

# By default, Go programs run with GOMAXPROCS set to the number of cores available
go run *.go -project gcp-expert-sandbox-jim -bucket gcp-expert-sandbox-jim-data-pipe -instance data-pipe -database data-pipe load
