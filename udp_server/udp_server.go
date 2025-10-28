package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Resolve the UDP address to listen on port 8081
	addr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Listen for incoming UDP packets
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server listening on port 8081...")

	buffer := make([]byte, 1024)

	// Loop forever to handle incoming packets
	for {
		// Read a UDP packet and get the client's address
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		message := strings.TrimSpace(string(buffer[:n]))
		fmt.Printf("Received message '%s' from %s\n", message, clientAddr)

		// Prepare the response
		response := "OK: " + message + "\n"

		// Send the response back to the client
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Error writing to UDP:", err)
		}
	}
}