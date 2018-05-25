
GOPATH:=$(shell go env GOPATH)

.PHONY: test docker run build


build:
	go build -o listen

run: build
	MICRO_SERVER_ADDRESS=:8000 ./listen

test:
	go test -v -cover -tags test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

docker:
	docker build . -t registry.gitlab.com/gitmate-micro/listen:latest

run-docker:
	docker run -d -p 8000:8000 registry.gitlab.com/gitmate-micro/listen:latest
