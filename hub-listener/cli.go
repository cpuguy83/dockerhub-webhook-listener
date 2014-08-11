package main

import (
	"flag"

	server "github.com/cpuguy83/dockerhub-webhook-listener"
)

var listenAddr = flag.String("listen", "localhost:8080", "<address>:<port> to listen on")

func main() {
	flag.Parse()

	server.Serve(*listenAddr)
}
