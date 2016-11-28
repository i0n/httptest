#!/bin/sh

# This script will compile OS X and Linux binaries in a docker container and place them in /bin.

# stop execution of script on failure
set -e

current_dir=`pwd`
script_rel_dir=`dirname $0`
script_dir=`cd $script_rel_dir && pwd`
root_dir=`dirname $script_dir`

# Load settings
. "$script_dir/settings.sh"

docker run \
  -v $root_dir:/go/src/$REPO_PROVIDER/$ORGANISATION/$PROJECT_NAME \
  -e ORGANISATION=$ORGANISATION \
  -e PROJECT_NAME=$PROJECT_NAME \
  -e REPO_PROVIDER=$REPO_PROVIDER \
  -e PROJECT_PATH=$PROJECT_PATH \
  -e BUILD_OSX=$BUILD_OSX \
  -e BUILD_LINUX=$BUILD_LINUX \
  $DOCKER_IMAGE \
  /bin/make.sh
