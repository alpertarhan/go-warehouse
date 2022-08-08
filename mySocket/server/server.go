package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Server() {
	// # open connection to server socket on port 8080
	// # and send the message "Hello World"
	// # close the connection

	fmt.Println("Starting server...")

	// listen on port 8080
	ln, _ := net.Listen("tcp", "8080")

	// accept connection from client
	conn, _ := ln.Accept()

	// handle multiple clients
	for {
		// listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// output message received
		fmt.Print("Message Received:", string(message))

		// sample process for string received
		newmessage := strings.ToUpper(message)

		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}

}
