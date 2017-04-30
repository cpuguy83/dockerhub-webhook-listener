package main

import (
	"flag"
	"log"

	"gopkg.in/gcfg.v1"

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
	if err := server.Serve(config); err != nil {
		log.Fatal(err)
	}
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
