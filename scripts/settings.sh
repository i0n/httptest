#!/bin/sh

# This script is for setting options used by all of the other scripts.

# stop execution of script on failure
set -e

ORGANISATION=i0n
PROJECT_NAME=httptest
REPO_PROVIDER=github.com
PROJECT_PATH=./cmd/cli/main.go
DOCKER_IMAGE=i0nw/go-and-node-dev:latest

# Do not edit below
if [ -z "$ORGANISATION" ]; then
  echo "ORGANISATION not set"
  echo "! Check scripts/settings.sh. Aborting..."
  exit 1
fi

if [ -z "$PROJECT_NAME" ]; then
  echo "PROJECT_NAME not set"
  echo "! Check scripts/settings.sh. Aborting..."
  exit 1
fi

if [ -z "$REPO_PROVIDER" ]; then
  echo "REPO_PROVIDER not set"
  echo "! Check scripts/settings.sh. Aborting..."
  exit 1
fi

if [ -z "$DOCKER_IMAGE" ]; then
  echo "DOCKER_IMAGE not set"
  echo "! Check scripts/settings.sh. Aborting..."
  exit 1
fi
