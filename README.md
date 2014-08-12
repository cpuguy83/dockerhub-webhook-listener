# DockerHub Webhook Listener

This is just a simple HTTP server setup to listen for and handle DockerHub
webhook calls.

A simple `Logger` handler and a slightly more complex `Mailgun` handler are
included for reference in creating handlers.

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

You should use SSL.
To do so ad a `tls` section to your config file, with a `cert` and a `key` file

You should also use authentication.
Right now DockerHub doesn't really support this, but you can use an api key as a
query param.
To handle this, you need to add an `apikeys` section to the config file along
with a list of `key`'s
