package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}
	g := gzip.NewWriter(w)
	defer func() {
		if err := g.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	multiWriter := io.MultiWriter(os.Stdout, g)

	// Write json data to both of gzip writer and standard output.
	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(source); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
