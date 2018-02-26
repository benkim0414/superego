#!/bin/bash

SUPEREGO_ROOT="$(dirname "${BASH_SOURCE}")/.."

readonly DOCKER_HUB_REPO=benkim
SUPEREGO_BUILD_IMAGE=${1:-"$DOCKER_HUB_REPO/superego"}
SUPEREGO_BUILD_IMAGE_TAG=${2:-"dev"}

docker build -t $SUPEREGO_BUILD_IMAGE:$SUPEREGO_BUILD_IMAGE_TAG -f $SUPEREGO_ROOT/hack/docker/Dockerfile $SUPEREGO_ROOT
