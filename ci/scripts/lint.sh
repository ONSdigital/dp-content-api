#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-content-api
  make lint
popd
