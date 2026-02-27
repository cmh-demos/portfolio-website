#!/bin/bash
# Helper script to run Terraform using the official image, avoiding host
# installation.  Usage: ./run.sh init|apply|plan [args]

docker run --rm -it \
  -v "$PWD":/workspace \
  -w /workspace \
  hashicorp/terraform:latest "$@"
