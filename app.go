package uzgo

import (
	"log"
	"net"
	"fmt"
)

func main() {

	// 启动
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

func broadcaster() {

}

func handleConn(conn) {
	fmt.Println(conn)
}
