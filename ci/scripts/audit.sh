#!/bin/bash -eux

export cwd=$(pwd)

pushd $cwd/dp-content-api
  make audit
popd