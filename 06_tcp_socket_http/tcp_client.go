package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	sendMessages := []string{"ASCII", "PROGRAMMING", "PLUS"}
	current := 0
	var conn net.Conn = nil

	for {
		var err error
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8080")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		request, err := http.NewRequest(
			"POST",
			"http://localhost:8080",
			strings.NewReader(sendMessages[current]),
		)
		if err != nil {
			log.Fatal(err)
		}

		err = request.Write(conn)
		if err != nil {
			log.Fatal(err)
			conn = nil
			continue
		}

		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			log.Fatal(err)
		}

		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(dump))
		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
