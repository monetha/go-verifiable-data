#!/bin/bash


set -e -o pipefail

if [ -z "$CI_PIPELINE_ID" ]
  then
    echo "No CI_PIPELINE_ID supplied."
    exit 1
fi

if [ -z "$PACKAGE_NAME" ]
  then
    echo "No PACKAGE_NAME supplied."
    exit 1
fi

GOLANG_VERSION=1.11
PACKAGE_FULL_PATH=/go/src/$PACKAGE_NAME
VERSION=$(date +"%Y.%m.%d").$CI_PIPELINE_ID

docker run -i --rm \
-v "$PWD":$PACKAGE_FULL_PATH \
-w $PACKAGE_FULL_PATH \
golang:$GOLANG_VERSION /bin/bash << COMMANDS
set -e -o pipefail
HOME=$PACKAGE_FULL_PATH
make dependencies
make lint
make test
COMMANDS

echo Done.
