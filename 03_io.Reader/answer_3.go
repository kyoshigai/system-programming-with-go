package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	archive, err := os.Create("archive.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	w, err := zipWriter.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(w, strings.NewReader("hoge"))
}
