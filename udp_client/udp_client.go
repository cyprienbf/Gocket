package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Resolve the server address
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}

	// Connect to the UDP server (sets a default destination for packets)
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Read input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')

	// Send the message to the server
	_, err = conn.Write([]byte(text))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// Wait for the response
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Response from server: %s", string(buffer[:n]))
}
