# DockerHub Webhook Listener

This is just a simple HTTP server setup to listen for and handle DockerHub
webhook calls.


## Usage

```bash
go run * -listen :8080
```

Register a handler in `handler.go`
You an use "Logger" as a reference to how to set this up.

## TODO
1. Implement api key so we can know it's the hub making the call
2. Use TLS, because TLS
