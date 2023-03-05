package main

import (
	"io"
	"log"
	"os"
)

func main() {
	inFileName := os.Args[1]
	outFileName := os.Args[2]

	inFile, err := os.Open(inFileName)
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(outFile, inFile)
}
