# Listen Service

This is the Listen service

Generated with

```
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

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./listen-web
```

Build a docker image
```
make docker
```
