
GOPATH:=$(shell go env GOPATH)

.PHONY: test docker run build


build:
	go build -o listen

run: build
	MICRO_SERVER_ADDRESS=:8000 ./listen

test:
	go test -v ./... -cover

docker:
	docker build . -t registry.gitlab.com/nkprince007/listen:latest
