package main

import (
	"flag"
	"log"

	"code.google.com/p/gcfg"

	server "github.com/cpuguy83/dockerhub-webhook-listener"
)

var listenAddr = flag.String("listen", "localhost:8080", "<address>:<port> to listen on")
var configFile = flag.String("config-file", "", "Location of handler config file")

func main() {
	flag.Parse()

	config, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on %s", *listenAddr)
	server.Serve(config)
}

func parseConfig() (*server.Config, error) {
	config := &server.Config{}
	if *configFile != "" {
		err := gcfg.ReadFileInto(config, *configFile)
		if err != nil {
			return nil, err
		}
	}

	config.ListenAddr = *listenAddr

	return config, nil
}
