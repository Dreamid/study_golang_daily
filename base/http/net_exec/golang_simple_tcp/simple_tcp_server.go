package golang_simple_tcp

import (
	"flag"
	"fmt"
	"net"
)

const maxRead = 25

func main() {
	flag.Parsed()
	if flag.NArg() != 2 {
		panic("usage：host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	listener.Accept()

}

func initServer(hostAndPort string) *net.TCPListener {

	return nil
}
