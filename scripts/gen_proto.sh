#!/bin/bash

CLIENT_PACKAGE=./pkg/trtclient/nvidia_inferenceserver
mkdir -p ${CLIENT_PACKAGE}
protoc -I ./protos --go_out=plugins=grpc:${CLIENT_PACKAGE} ./protos/*.proto
