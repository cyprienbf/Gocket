#!/bin/bash

# This script simulates multiple TCP clients using netcat (nc).
echo "Starting 3 clients using netcat..."

# Function to run a client in the background
run_nc_client() {
  # Send a message to the server on localhost:8080 and print the response.
  (sleep $2; echo "Message from nc Client $1") | nc localhost 8080 &
}

# Check if the server is running on port 8080
if ! nc -z localhost 8080; then
    echo "Server is not running on port 8080. Please start server.go first."
    exit 1
fi

run_nc_client 1 0.1
run_nc_client 2 0.2
run_nc_client 3 0.3

echo "Clients launched in the background. Check server output and client responses."
wait
echo "All clients have finished."