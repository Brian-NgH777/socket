package client

import (
	"fmt"
	//"io/ioutil"
	"net"
	"os"
)

const (
	hostPythonServer = "127.0.0.1"
	portPythonServer = "5566"
)

func Client(input string) string {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: %s host:port ", os.Args)
	//     fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
	//     os.Exit(1)
	// }
	service := fmt.Sprintf("%s:%s",hostPythonServer,portPythonServer)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte(input))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	//result, err := ioutil.ReadAll(conn)
	//checkError(err)
	fmt.Println("result", string(buf[0:n]))
	return string(buf[0:n])
	//fmt.Println("resultresultresultresultresult", string(result))
	//os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
