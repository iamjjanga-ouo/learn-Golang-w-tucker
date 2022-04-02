package main

import (
	"fmt"
	"os"
)

type Writer func(str string)

func writeHello(writer Writer) {
	writer("Hello world")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) { 	//	1. writeHello 함수 호출
		fmt.Fprintln(f, msg)		//	2. literal 함수 인자로 전달
	})								//	3. msg : "Hello world" -> Fprintln(f, msg)로 파일 쓰기
}
