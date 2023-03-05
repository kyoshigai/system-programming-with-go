package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"Apple", "The United States", "Steve Jobs"},
		{"Nintendo", "Japan", "Fusajiro Yamauchi"},
	}
	file, err := os.Create("companies.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}
}
