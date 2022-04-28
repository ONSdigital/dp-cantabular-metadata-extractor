#!/bin/bash -eux

pushd dp-cantabular-metadata-extractor
  make test-component
popd
