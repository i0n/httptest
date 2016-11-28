#!/bin/sh

# This script will run all node tests in a docker container.
# pass no-docker-compose as an argument to avoid using docker-compose inside the script (workaround for CircleCI dodgy docker implementation)

# stop execution of script on failure
set -e

current_dir=`pwd`
script_rel_dir=`dirname $0`
script_dir=`cd $script_rel_dir && pwd`
root_dir=`dirname $script_dir`

# Load settings
. "$script_dir/settings.sh"

cd $root_dir

if [ "$1" = no-docker-compose ]; then
  echo "Running script without docker-compose commands..."
else
  docker-compose down
  docker-compose pull
  docker-compose up -d --force-recreate
fi

docker run \
  -v $root_dir:/go/src/$REPO_PROVIDER/$ORGANISATION/$PROJECT_NAME \
  -e ORGANISATION=$ORGANISATION \
  -e PROJECT_NAME=$PROJECT_NAME \
  -e REPO_PROVIDER=$REPO_PROVIDER \
  -e PROJECT_PATH=$PROJECT_PATH \
  $DOCKER_IMAGE \
  dockerize -timeout 120s go test -race -v $(go list ./... | grep -v /vendor/)
