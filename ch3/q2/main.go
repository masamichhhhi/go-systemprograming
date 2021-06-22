package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.Create("ch3/q2/random.txt")
	if err != nil {
		panic(err)
	}
	_, err = io.CopyN(file, rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
}
