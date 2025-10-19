package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	// Listen on TCP port 8080 on all available interfaces.
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	// Close the listener when the application closes.
	defer ln.Close()

	for {
		// Listen for an incoming connection.
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}

		// Handle connections in a new goroutine.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Close the connection when this function ends.
	defer conn.Close()

	// Create a new reader to read data from the connection.
	reader := bufio.NewReader(conn)

	// Read the incoming data until a newline is received.
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// Trim the newline character from the message.
	message = strings.TrimSpace(string(message))
	fmt.Printf("Message Received: %s\n", message)

	// Send a response back to the client.
	response := fmt.Sprintf("OK. You sent: %s\n", message)
	conn.Write([]byte(response))
}
