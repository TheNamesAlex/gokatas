package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ip := os.Args[1]
	log.Print("using ip: " + string(ip))
	conn, err := net.Dial("tcp", ip+":80")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {

	}
}
