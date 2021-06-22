package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("File Info")
	fmt.Printf("ファイル名：%v", info.Name())
	fmt.Printf("サイズ：%v", info.Size())
	fmt.Printf("変更日時：%v", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("ディレクトリ？: %v", info.Mode().IsDir())
	fmt.Printf("読み込み可能か？：%v", info.Mode().IsRegular())
	fmt.Printf("Unixのアクセス権限：%v", info.Mode().Perm())
	fmt.Printf("モード %v\n", info.Mode().String())
}
