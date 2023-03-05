package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	outFileName := os.Args[1]
	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal(err)
	}

	io.CopyN(outFile, rand.Reader, 1024)
}
