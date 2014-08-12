# DockerHub Webhook Listener

This is just a simple HTTP server setup to listen for and handle DockerHub
webhook calls.


## Usage
Add a config file using `-config-file`
This file should be in INI format and is intended for use with handlers
Register a handler in `handler.go`
You an use "Logger" as a reference to how to set this up.

```bash
cd hub-listener
go build
./hub-listener -listen 0.0.0.0:80 -config-file config.ini
```

## TODO
1. Implement api key so we can know it's the hub making the call
2. Use TLS, because TLS
