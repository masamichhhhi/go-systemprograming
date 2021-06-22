package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func writeConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	for sessionResponse := range sessionResponses {
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}

func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))
	content := "hello world\n"

	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          ioutil.NopCloser(strings.NewReader(content)),
	}
	resultReceiver <- response
}

//コネクションの数だけWriteConnでconnをつくって、あとからResponseを入れ込んでいくみたいなイメージ？
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)
	go writeConn(sessionResponses, conn)

	reader := bufio.NewReader(conn)
	for {
		conn.SetDeadline(time.Now().Add(5 * time.Second))

		request, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse
		go handleRequest(request, sessionResponse)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}

	// listener, err := net.Listen("tcp", "localhost:8888")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Server is running at localhost:8888")
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	go func() {
	// 		defer conn.Close()
	// 		fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// 		for {
	// 			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	// 			request, err := http.ReadRequest(bufio.NewReader(conn))
	// 			if err != nil {
	// 				neterr, ok := err.(net.Error)
	// 				if ok && neterr.Timeout() {
	// 					fmt.Println("Timeout")
	// 				} else if err == io.EOF {
	// 					break
	// 				}
	// 				panic(err)
	// 			}

	// 			dump, err := httputil.DumpRequest(request, true)
	// 			if err != nil {
	// 				panic(err)
	// 			}
	// 			fmt.Println(string(dump))
	// 			content := "hello world\n"

	// 			response := http.Response{
	// 				StatusCode:    200,
	// 				ProtoMajor:    1,
	// 				ProtoMinor:    1,
	// 				ContentLength: int64(len(content)),
	// 				Body:          ioutil.NopCloser(strings.NewReader(content)),
	// 			}
	// 			response.Write(conn)
	// 		}
	// 	}()
	// }
}
