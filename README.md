# listen

[![pipeline status](https://gitlab.com/gitmate-micro/listen/badges/master/pipeline.svg)](https://gitlab.com/gitmate-micro/listen/commits/master)

A simple go-micro web service that listens to incoming webhooks and sends out
events over specified topics.

Auto-generated with

```sh
micro new gitlab.com/gitmate-micro/listen --namespace=gitmate.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: gitmate.micro.web.listen
- Type: web
- Alias: listen

## Dependencies

Micro services depend on service discovery. The default is consul.

```sh
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```sh
make build
```

Run the service

```sh
./listen
```

Build a docker image

```sh
make docker
```

### With docker compose

Build the docker-compose network

```sh
docker-compose up -d
```

## Images / Container Registry

For released images, please visit the GitLab Container Registy for this
repository [here](https://gitlab.com/gitmate-micro/listen/container_registry).

To run the latest image, pull it with docker

```sh
# login to gitlab container registry with your credentials
docker login registry.gitlab.com

# run the container locally exposing port 8000
make run-docker
```
