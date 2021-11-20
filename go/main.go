package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
	buf := make([]byte, 1050)

	// Only read once
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	tmp := string(buf[:reqLen])
	index := strings.Index(tmp, "\r\n")
	fmt.Println(reqLen)
	fmt.Println(index)

	// url := string(buf[:reqLen])
	// for _, s := range buf[:reqLen] {

	// 	if string(s) == "A" {
	// 		return "A"
	// 	} else if string(s) == "B" {
	// 		return "B"
	// 	}
	// }

	// conn.Write([]byte("Message received."))
	conn.Write([]byte(buf[:reqLen]))
	conn.Write([]byte(tmp[:index]))

	conn.Close()
}
