#!/bin/bash -eux

pushd dp-content-api
  make build
  cp build/dp-content-api Dockerfile.concourse ../build
popd
