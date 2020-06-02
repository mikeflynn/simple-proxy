package main

import (
	"flag"
	"io"
	"log"
	"net"
)

func handler(src net.Conn, host string) {
	dst, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatalln("Unable to connect to the host.")
	}

	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	remote := flag.String("remote", "", "The remote address.")
	local := flag.String("local", "0.0.0.0:80", "The local listening address")
	flag.Parse()

	listener, err := net.Listen("tcp", *local)
	if err != nil {
		log.Fatalln("Unable to bind to port.")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handler(conn, *remote)
	}
}
