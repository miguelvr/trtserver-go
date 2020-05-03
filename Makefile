.PHONY: up predict clean

default: predict

proto:
	bash $(PWD)/scripts/gen_proto.sh

build:
	python $(PWD)/scripts/trace.py
	go build -o predict cmd/predict/main.go

clean:
	rm -rf $(PWD)/assets/models
	rm ./predict

up:
	docker run --rm -d -it \
		--ipc host \
		--name trtserver \
		-p 8000:8000 -p 8001:8001 -p 8002:8002 \
		-v $(PWD)/assets/models:/models \
		nvcr.io/nvidia/tritonserver:20.03-py3 \
		trtserver --model-repository /models

down:
	docker stop trtserver

predict: up build
