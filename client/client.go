package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server.
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	for {
		// Read input from the user.
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// Send the text to the server.
		fmt.Fprintf(conn, text)

		// Read the response from the server.
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err.Error())
			return
		}
		fmt.Print("Message from server: " + message)
	}
}
