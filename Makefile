
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


build:
	go build -o listen

test:
	go test -v ./... -cover

docker:
	docker build . -t listen-web:latest
