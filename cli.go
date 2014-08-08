package main

import "flag"

var listenAddr = flag.String("listen", "localhost:8080", "<address>:<port> to listen on")

func main() {
	flag.Parse()

	Serve(*listenAddr)
}
