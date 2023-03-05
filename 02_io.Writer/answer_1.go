package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(file, "int: %d, float: %f, string: %s", 1, 1.1, "hoge")
}
