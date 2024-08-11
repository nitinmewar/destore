package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// listen for TCP connections
	fmt.Println("listening on port http://localhost:6379")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error while reading from client : ", err.Error())
			os.Exit(1)
		}

		fmt.Println("mesage recieved : ", string(buf))

		conn.Write([]byte("+OK\r\n"))
	}
}
