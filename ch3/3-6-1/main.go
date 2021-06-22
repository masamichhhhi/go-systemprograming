package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `a
b
c`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%v#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
