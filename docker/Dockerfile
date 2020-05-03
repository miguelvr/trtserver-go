FROM python:3.7 as model_builder

WORKDIR /opt/trtserver-go

COPY ./assets ./assets
COPY ./scripts ./scripts

RUN pip install torch==1.4.0+cpu torchvision==0.5.0+cpu -f https://download.pytorch.org/whl/torch_stable.html

RUN python scripts/trace.py

##################################################

FROM golang:1.14-alpine3.11 as go_builder

COPY --from=model_builder /opt/trtserver-go/assets/models/ /opt/trtserver-go/assets/models/

WORKDIR /opt/trtserver-go
COPY . .

ENV CGO_ENABLED 0
ENV GOOS linux
RUN go mod download
RUN go build -o predict cmd/predict/main.go

##################################################

FROM alpine:3.11

COPY --from=go_builder /opt/trtserver-go/predict ./
