#!/bin/bash

# With -v, this will let the host share its ADC to the container.
docker run --env-file config.env -it --rm -v $HOME/.config:/root/.config gcr.io/$(gcloud config get-value core/project)/data-pipe:v1 generate
