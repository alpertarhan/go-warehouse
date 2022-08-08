// # simple socket server example

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8080")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {

		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// output message received
		fmt.Print("Message Received:", string(message))

		// sample process for string received
		newmessage := strings.ToUpper(message)

		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
		time.Sleep(200 * time.Millisecond)
	}
}

// # simple socket client example

func client() {

	// connect to this socket
	conn, _ := net.Dial("tcp", ":8080")

	// send a message to the socket
	fmt.Fprintf(conn, "Hello from the client")

	// listen for reply

	message, _ := bufio.NewReader(conn).ReadString('\n')

	// output message received
	fmt.Print("Message Received:", string(message))

	os.Exit(0)

}
