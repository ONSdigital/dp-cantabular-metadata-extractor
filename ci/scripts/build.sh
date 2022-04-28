#!/bin/bash -eux

pushd dp-cantabular-metadata-extractor
  make build
  cp build/dp-cantabular-metadata-extractor Dockerfile.concourse ../build
popd
