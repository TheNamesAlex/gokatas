package main

import (
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// Send a message to the client every second
		_, err := c.Write([]byte("Hello from server at: " + time.Now().Format("15:04:05") + "\n"))
		if err != nil {
			log.Print("Error writing to connection: ", err)
			return
		}
		time.Sleep(1 * time.Second) // Wait for a second before sending the next message
	}
}
