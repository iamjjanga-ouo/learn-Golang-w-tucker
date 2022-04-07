package main

import "fmt"

// 채널에서 데이터를 안가져가는 경우
func main() {
	ch := make(chan int)

	ch <- 9
	fmt.Println("Never print")

}
