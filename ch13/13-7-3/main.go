package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("初期化")
}

var once sync.Once

func main() {
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
