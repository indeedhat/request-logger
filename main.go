package main

import (
	"log"
	"net"
	"os"
	"sync"
)

var response = `HTTP/1.1 200 OK

`

var reqMux sync.Mutex

func main() {
	fh, err := os.OpenFile("request.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	log.SetOutput(fh)

	soc, err := net.Listen("tcp", ":8084")
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

	reqMux.Lock()
	defer reqMux.Unlock()

	log.Println(string(buf))

	con.Write([]byte(response))
}
