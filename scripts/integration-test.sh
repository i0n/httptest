#!/bin/sh

# This script will run all integartion tests in a docker container.
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
  docker pull i0nw/httptest:latest 
  docker-compose up -d --force-recreate
fi

docker run -ti \
  --net httptest \
  --link exampleServer_dev_local:exampleServer \
  -v $root_dir:/opt/$PROJECT_NAME \
  -e CONFIG_FILE=/opt/$PROJECT_NAME/test/fixtures/fixtures.json \
  i0nw/httptest:latest \
  dockerize -wait http://exampleServer:8080/health -timeout 120s httptest

