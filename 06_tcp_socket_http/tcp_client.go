package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
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
		request.Header.Set("Accept-Encoding", "gzip")

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

		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(dump))

		defer response.Body.Close()
		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
