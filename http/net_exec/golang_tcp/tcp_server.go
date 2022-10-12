package golang_tcp

import (
	"fmt"
	"net"
)

func StartServer() {

	fmt.Println("Starting the tcp server ...")
	//	创建listener
	listner, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	//	监听并接收来自客户端的连接
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)

	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf[:len]))

	}
}
