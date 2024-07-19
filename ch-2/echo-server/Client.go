package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server.
	conn, err := net.Dial("tcp", "localhost:20080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Send a message to the server.
	message := "Hello, Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to server:", err)
		return
	}

	// Buffer to store response from server.
	buf := make([]byte, 512)
	// Read the response from server.
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}

	// Print the response from server.
	fmt.Println("Received from server:", string(buf[:n]))
}
