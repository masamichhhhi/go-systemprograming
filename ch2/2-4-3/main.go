package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("byte.Buffer example\n"))
	fmt.Println(buffer.String())
}
