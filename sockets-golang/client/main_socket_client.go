package main

import (
	"fmt"
	//"io/ioutil"
	"net"
	"os"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: %s host:port ", os.Args)
	//     fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
	//     os.Exit(1)
	// }
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	for i := 1; i <= 5; i++ {
		str := fmt.Sprintf("nguyen huu thi %d", i)
		_, err = conn.Write([]byte(str))
		checkError(err)
	}
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	//result, err := ioutil.ReadAll(conn)
	//checkError(err)
	fmt.Println("resultresultresultresultresult", string(buf[0:n]))
	//fmt.Println("resultresultresultresultresult", string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
