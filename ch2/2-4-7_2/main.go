package main

import (
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダー")
	request.Write(os.Stdout)
}
