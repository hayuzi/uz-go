// netcat 是一个只读的TCP客户端程序
package client

import (
	"io"
	"net"
	"os"
	"fmt"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("lack of args")
	}

	host := args[0]
	port, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("port error")
	}

	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(address)
		fmt.Println("address connect error")
		// log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		// log.Fatal(err)
		fmt.Println("src string copy error")
	}
}
