package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "1966"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
		fmt.Println("net.Listen err=", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	fmt.Printf("server listening %s:%s", CONN_HOST, CONN_PORT)
	fmt.Println()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("net.Listen.Accept", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	// Only read once
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Printf("%s", buf[:reqLen])

	// conn.Write([]byte("Message received."))
	conn.Write([]byte(buf[:reqLen]))

	conn.Close()
}
