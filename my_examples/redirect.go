package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// handleConnection handles the incoming connection and redirects data
func handleConnection(conn net.Conn) {

	// Start two goroutines to handle bi-directional data transfer
	go func() {
		rp, wp := io.Pipe()
		go func() {
			io.Copy(wp, conn)
		}()
		io.Copy(conn, rp)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "Error copying from destination to source: %v\n", err)
		//}
	}()
	time.Sleep(3 * time.Second)
}

func main() {

	localPort := "8082"

	localAddress := fmt.Sprintf(":%s", localPort)

	listener, err := net.Listen("tcp", localAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting TCP server: %v\n", err)
		os.Exit(1)
	}
	//defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accepting connection: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}
}
