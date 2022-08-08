package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client() {

	// connect to the server
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")

	// what to send
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output
		fmt.Print("Message from server: " + message)

	}

}
