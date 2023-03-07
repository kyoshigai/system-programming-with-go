package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running at localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			for {
				request, err := http.ReadRequest(
					bufio.NewReader(conn),
				)
				if err != nil {
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					log.Fatal(err)
				}
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(string(dump))
				content := "Hello world\n"
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          io.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}
