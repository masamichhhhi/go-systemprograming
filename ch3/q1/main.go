package main

import (
	"flag"
	"io"
	"os"
)

var (
	stringOpt = flag.String("s", "[-s オプションで追加した文字列だよ!]", "help message for b option")
)

func main() {

	src := os.Args[1]
	dest := os.Args[2]

	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, _ := os.Create(dest)
	defer newFile.Close()
	io.Copy(newFile, file)

	flag.Parse()
	newFile.Write([]byte(*stringOpt))
}
