#!/bin/sh

# This script will resolve node dependencies for the project and place them in node_modules/

# stop execution of script on failure
set -e

current_dir=`pwd`
script_rel_dir=`dirname $0`
script_dir=`cd $script_rel_dir && pwd`
root_dir=`dirname $script_dir`

# Load settings
. "$script_dir/settings.sh"

# If additional args passed, proxy to glide, otherwise default to install
if [ "$1" != "" ]; then
  COMMAND_ARGS="$@"
else
  COMMAND_ARGS="install"
fi

echo "Running -->"
echo "cd /go/src/$REPO_PROVIDER/$ORGANISATION/$PROJECT_NAME && glide '$COMMAND_ARGS'"

docker run \
  -ti \
  --rm \
  -v $root_dir:/go/src/$REPO_PROVIDER/$ORGANISATION/$PROJECT_NAME \
  -e PROJECT_NAME=$PROJECT_NAME \
  $DOCKER_IMAGE \
  sh -c "cd /go/src/$REPO_PROVIDER/$ORGANISATION/$PROJECT_NAME && glide $COMMAND_ARGS"
