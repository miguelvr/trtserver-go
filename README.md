# Serving a PyTorch model with Triton Inference Server and GoLang

## Introduction

This repository demonstrates how to serve a Pytorch model with 
NVIDIA's Triton Inference Server (TRT Server) and a Go client.

The Pytorch model is served by the TRT Server using libtorch and C++, 
which can only handle binary inputs and outputs. 
The goal of the Go client is to pre-process the input, encode the inputs and 
decode the outputs from the communication with the TRT Server.

An example is provided by serving a torchvision ResNet-50 model pre-trained on ImageNet.

## Requirements

- Docker
- GoLang 1.14
- Python 3.6+
- Nvidia Docker (optional, for GPUs)

## Getting Started

Trace the ResNet-50 model and build the corresponding Go Triton Client.
```shell script
make build
``` 

Start the docker container
```shell script
make up
```

You can now send a prediction with the Go client to the Triton Server
```shell script
./predict assets/white_shark.jpg
```
```
# expected output
TRTIS Health - Live: true
TRTIS Status: map[resnet50_imagenet:config:{name:"resnet50_imagenet"  platform:"pytorch_libtorch"  version_policy:{latest:{num_versions:1}}  max_batch_size:32  input:{name:"INPUT__0"  data_type:TYPE_FP32  format:FORMAT_NCHW  dims:3  dims:224  dims:224}  output:{name:"OUTPUT__0"  data_type:TYPE_INT32  dims:1  label_filename:"imagenet_labels.txt"}  output:{name:"OUTPUT__1"  data_type:TYPE_FP32  dims:1}  optimization:{input_pinned_memory:{enable:true}  output_pinned_memory:{enable:true}}  instance_group:{name:"resnet50_imagenet"  kind:KIND_CPU  count:1}  default_model_filename:"model.pt"}  version_status:{key:1  value:{ready_state:MODEL_READY  ready_state_reason:{}  infer_stats:{key:1  value:{success:{count:1  total_time_ns:196770676}  compute:{count:1  total_time_ns:196686936}  queue:{count:1  total_time_ns:73150}}}  model_execution_count:1  model_inference_count:1  last_inference_timestamp_milliseconds:14466241146505000000}}]
Inference Response: {"label": WHITE SHARK, "label_id": 2, "score": 0.983945}
```

Whenever you are done, take down the running triton inference server run the following:
```shell script
make down
```

To remove the created model artifacts and binaries by the build command run:
```shell script
make clean
```

## Metrics

You can access the server metrics by accessing to http://localhost:8002/metrics
