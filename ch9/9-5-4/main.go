package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))

	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(abspath)

	relpath, _ := filepath.Rel("/usr/local/go/src/", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relpath)
}
