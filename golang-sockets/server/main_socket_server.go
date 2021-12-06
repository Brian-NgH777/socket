package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	host = "127.0.0.1"
	port = "65432"
)

func main() {
	// port := os.Getenv("PORT")
	//port := "65432"
	service := fmt.Sprintf("%s:%s",host, port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("tcpAddr:", tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Println("Socket run port:", service)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// run as a goroutine
		go handleClient(conn)

		// handleClient(conn)
        // conn.Close() // we're finished
	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()

	var buf [512]byte
	for {
		// read upto 512 bytes
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		fmt.Println("Message and RemoteAddr", string(buf[0:n]), conn.RemoteAddr())
		daytime := time.Now()
		reply := fmt.Sprintf("%s_%s",string(buf[0:n]), daytime.Format("2006-01-02 15:04:05"))
		// write the n bytes read
		_, err2 := conn.Write([]byte(reply))
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}