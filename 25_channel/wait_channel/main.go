package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // 채널이 닫히면 for loop 탈출
		fmt.Printf("Square : %d\n", n*n)
		// time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 채널 닫음
	wg.Wait()
}
