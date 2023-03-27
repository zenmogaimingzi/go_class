package main

import (
	"fmt"
	"net"
)

func main() {

	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8888")
	listen, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	for {
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)

		// 断开连接之后会报错提示 防止循环一直报错需要break 退出
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String() + "===" + string(buff[0:n]))
		conn.Write([]byte("收到了" + string(buff[0:n])))

	}

}
