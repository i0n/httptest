#!/bin/sh

# This script checks if you are on master branch and exits if not.

set -e

ORIGINAL_DIR=$(pwd)

if [ -z "$1" ]; then
  echo "You did not pass a directory as first arg, using current directory." 
else
  cd $1
fi

cleanup() {
  cd $ORIGINAL_DIR
}
# Always run cleanup...
trap cleanup INT TERM EXIT

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

LATEST_TAG=$(git describe --abbrev=0 --tags)
TAG_COMMIT=$(git rev-parse ${LATEST_TAG}~0)
CURRENT_MASTER_COMMIT=$(git rev-parse master)

echo "Checking to see if last master commit matches last tag commit."
echo
echo "LATEST TAG:            $LATEST_TAG"
echo "TAG COMMIT:            $TAG_COMMIT"
echo "CURRENT MASTER COMMIT: $CURRENT_MASTER_COMMIT"
echo

if [ "$CURRENT_MASTER_COMMIT" != "$TAG_COMMIT" ]; then
  echo "Tag does not match last commit to master. Exiting..." 
  exit 1
fi
