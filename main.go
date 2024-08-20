package main

import (
	"log"
	"net"
)

var response = `HTTP/1.1 200 OK

`

func main() {
	soc, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		con, err := soc.Accept()
		if err != nil {
			log.Printf("accept failed: %s", err)
			continue
		}

		go handle(con)
	}
}

func handle(con net.Conn) {
	defer con.Close()

	buf := make([]byte, 2048)

	_, err := con.Read(buf)
	if err != nil {
		log.Printf("read failed: %s", err)
		return
	}

	log.Println(string(buf))
	con.Write([]byte(response))
}
