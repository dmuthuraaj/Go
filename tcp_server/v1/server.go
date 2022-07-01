package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//closing listener
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(con net.Conn) {
	buffer := make([]byte, 1024)
	_, err := con.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	time := time.Now().Format(time.ANSIC)
	io.WriteString(con, "Hello\n")
	con.Write([]byte("Hi\n"))
	con.Write([]byte(time))
	con.Close()
}
