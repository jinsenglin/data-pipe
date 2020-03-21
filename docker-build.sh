#!/bin/bash

# This will copy all your dependencies from your GOPATH to the vendor directory.
govendor add +external

docker build -t gcr.io/$(gcloud config get-value core/project)/data-pipe:v1 --build-arg version=v1 .
