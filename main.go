package main

import (
	"Chat/server"
	"flag"
)

var (
	host *string
	port *int
)

func init() {
	host = flag.String("h", "localhost", "hostname")
	port = flag.Int("p", 3090, "port")
}

func main() {
	s := server.NewServer()
	s.LoadServerComponents()
	s.Initialize(host, port)
}
