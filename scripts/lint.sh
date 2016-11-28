#!/bin/sh

# This script will run all go tests in a docker container.

# stop execution of script on failure
set -e

current_dir=`pwd`
script_rel_dir=`dirname $0`
script_dir=`cd $script_rel_dir && pwd`
root_dir=`dirname $script_dir`

# Load settings
. "$script_dir/settings.sh"

cd $root_dir

docker-compose pull
docker-compose up -d --no-recreate

docker run \
  -v $root_dir:/go/src/$REPO_PROVIDER/$ORGANISATION/$PROJECT_NAME \
  -e ORGANISATION=$ORGANISATION \
  -e PROJECT_NAME=$PROJECT_NAME \
  -e REPO_PROVIDER=$REPO_PROVIDER \
  -e PROJECT_PATH=$PROJECT_PATH \
  $DOCKER_IMAGE \
  gometalinter --disable-all --skip=pb --skip=utils/copyfile --skip=utils/readfline --skip=utils/writefchunk --skip=utils/regenStubWithString.go --enable=gofmt --enable=vet --enable=vetshadow --disable=errcheck --enable=golint --concurrency=4 --vendor --deadline=5m --sort=path ./src/$REPO_PROVIDER/$ORGANISATION/$PROJECT_NAME
