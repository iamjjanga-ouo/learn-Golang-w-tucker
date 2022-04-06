package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d초", i)
		time.Sleep(1 * time.Second) // main 함수의 time.Sleep이 더 짧으면 다른 함수 진행도중 프로그램 종료가 먼저됨
	}

	fmt.Println("끝!")
}
