package main

import (
	"mysocket/client"
	"mysocket/server"
)

func main() {
	server.Server()
	client.Client()
}
