package golang_tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartClient() {
	//	打开连接、
	conn, err := net.Dial("tcp", "localhost:8888")

	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Println("First, what is your name?")
		clientName, _ := inputReader.ReadString('\n')
		trimmedClient := strings.Trim(clientName, "\r\n")
		fmt.Println("what to send to the server? Type Q to quit.")

		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
