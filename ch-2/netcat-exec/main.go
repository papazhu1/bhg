package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {

	fmt.Println("handle exec")
	/*
	 * Explicitly calling /bin/sh and using -i for interactive mode
	 * so that we can use it for stdin and stdout.
	 * For Windows use exec.Command("cmd.exe")
	 */
	// cmd := exec.Command("cmd.exe")
	// cmd := exec.Command("/bin/sh", "-i")

	// 这里rp, wp是管道返回的 reader 和 writer 的地址
	//rp, wp := io.Pipe()
	//// Set stdin to our connection
	//cmd.Stdin = conn
	//cmd.Stdout = wp

	// copy 的左边是 dst，右边是src
	// 等待cmd的输出，然后写入conn
	go func() {
		fmt.Println("begin copy")
		newconn, err := net.Dial("tcp", "localhost:8082")
		if err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println("dial newconn success")
		}
		// io.Copy(os.Stdout, conn)
		go func() {
			io.Copy(newconn, conn)
		}()
		// io.Copy(os.Stdout, newconn)
		// fmt.Println("copy from conn to newconn success")
		go func() {
			io.Copy(conn, newconn)
		}()
		fmt.Println("end copy")
	}()

}

func main() {
	listener, err := net.Listen("tcp", "localhost:20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		fmt.Println("accept success")
		fmt.Println("conn", conn)
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
