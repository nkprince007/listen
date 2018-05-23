# listen

[![pipeline status](https://gitlab.com/nkprince007/listen/badges/master/pipeline.svg)](https://gitlab.com/nkprince007/listen/commits/master)

A simple go-micro web service that listens to incoming webhooks and sends out
events over specified topics.

Auto-generated with

```sh
micro new gitlab.com/nkprince007/listen --namespace=go.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.listen
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
