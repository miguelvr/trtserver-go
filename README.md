# Serving a PyTorch model with Triton Inference Server and GoLang

## Introduction

This repository demonstrates how to serve a Pytorch model with 
NVIDIA's Triton Inference Server (TRT Server) and a Go client.

The Pytorch model is served by the TRT Server using libtorch and C++, 
which can only handle binary inputs and outputs. 
The goal of the Go client is to pre-processe the input, encode the inputs and 
decode the outputs from the communication with the TRT Server.

An example is provided by serving a torchvision ResNet-50 model pre-trained on ImageNet.

## Requirements

- Docker
- GoLang 1.14
- Python 3.6+
- Nvidia Docker (optional, for GPUs)

## Getting Started

Start the docker container
```shell script
make up
```

Trace the ResNet-50 model and build the corresponding Go Triton Client.
```shell script
make build
``` 

You can now send a prediction with the Go client to the Trtiton Server
```shell script
./predict
```

Whenever you are done, take down the running triton inference server run the following:
```shell script
make down
```

To remove the created model artifacts and binaries by the build command run:
```shell script
make clean
```
