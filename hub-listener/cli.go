package main

import (
	"flag"
	"log"

	server "github.com/cpuguy83/dockerhub-webhook-listener"
)

var listenAddr = flag.String("listen", "localhost:8080", "<address>:<port> to listen on")

func main() {
	flag.Parse()

	log.Printf("Starting server on %s", *listenAddr)
	server.Serve(*listenAddr)
}
