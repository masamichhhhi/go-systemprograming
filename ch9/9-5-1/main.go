package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Temp file Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))
}
