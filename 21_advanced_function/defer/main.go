package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("반드시 호출")
	defer f.Close()
	defer fmt.Println("파일 닫힘")

	fmt.Println("파일에 Hello world를 씀")
	fmt.Fprintln(f, "Hello World")
}
