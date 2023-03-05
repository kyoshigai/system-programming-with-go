package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	zipFile, err := zipWriter.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(zipFile, strings.NewReader("hogehogehogehogehoge"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
