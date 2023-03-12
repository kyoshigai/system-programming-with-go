package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/edsrzf/mmap-go"
)

func main() {
	// テストデータを書き込み
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := ioutil.WriteFile(testPath, testData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Unmap()

	m[9] = 'X'
	m.Flush()
	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap: %s\n", m)
	fmt.Printf("file: %s\n", fileData)
}
