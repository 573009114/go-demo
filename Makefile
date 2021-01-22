.PHONY: run build
build:
	GOOS=linux GOARCH=amd64 go build -tags netgo -o bin/test cmd/main.go
	docker build -f build/docker/Dockerfile -t test .
run:
	docker run --rm  -e MICRO_REGISTRY=mdns test

