package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.After(5 * time.Second)
L:
	for {
		select {
		case <-timer:
			fmt.Println("5 secs passed.")
			break L
		default:
			fmt.Println("Not yet")
		}
		time.Sleep(time.Second)
	}
}
