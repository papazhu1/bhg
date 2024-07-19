package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	for i := 1; i <= 10240; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	time.Sleep(5 * time.Second)
}
