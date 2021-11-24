#!/bin/bash

set -ex

repo="wellharbor.westwell-research.com/common"

docker build -t ${repo}/golang:1.17-buster $(dirname $0)/build/go
docker build -t ${repo}/helm:3.7.1 $(dirname $0)/build/helm

docker push ${repo}/golang:1.17-buster
docker push ${repo}/helm:3.7.1 