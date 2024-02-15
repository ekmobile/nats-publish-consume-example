# nats-publish-consume-example

This example is used to find a minimal example for a publisher / consumer structure based upon https://pkg.go.dev/github.com/nats-io/nats.go@v1.32.0/jetstream#readme-basic-usage with NATS running in a Docker container.

# Prerequisite

- Docker Desktop
- Golang

# Usage

start nats

```bash
docker-compose up
```

run the publisher

```bash
cd publisher
go get
go run .
```

run the consumer

```bash
cd consumer
go get
go run .
```

# Problem

the consumer does not display messages, although it should via the callback.

Root Cause: Main Program exits directly. Wait for it.
