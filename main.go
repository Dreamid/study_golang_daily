package main

import "golang_demo/golang_study/golang_net/golang_tcp"

func main() {
	go golang_tcp.StartServer()
	golang_tcp.StartClient()
}
