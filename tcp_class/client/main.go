package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("321")

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	reder := bufio.NewReader(os.Stdin)

	for {
		bytes, _, _ := reder.ReadLine()
		conn.Write(bytes)
		buff := make([]byte, 1024)
		num, _ := conn.Read(buff)
		fmt.Println(string(buff[0:num]))
	}

}
