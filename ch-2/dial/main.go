package main

import (
	"fmt"
	"net"
)

func main() {
	// _, err := net.Dial("tcp", "scanme.nmap.org:80")
	_, err := net.Dial("udp", "localhost:20080")
	if err == nil {
		fmt.Println("Connection successful")
	}
}
