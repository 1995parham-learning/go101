// tiny server in go for echoing back what you say
// in pure tcp
package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1373")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)

			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	if _, err := io.Copy(c, c); err != nil {
		return
	}
	defer c.Close()
}
