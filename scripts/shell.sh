#!/bin/sh

# This script will run put you in payfreindz_go container so you can have a look around.

# stop execution of script on failure
set -e

current_dir=`pwd`
script_rel_dir=`dirname $0`
script_dir=`cd $script_rel_dir && pwd`
root_dir=`dirname $script_dir`

# Load settings
. "$script_dir/settings.sh"

cd $root_dir

docker exec -ti i0n_${PROJECT_NAME}_dev_local sh
