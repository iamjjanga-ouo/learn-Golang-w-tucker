package main

import (
	"fmt"
	"sync"
	"time"
)

// select : 여러 채널에서 들어오기를 대기하는 상황 or 여러 채널을 동시에 대기하고 싶은경우
func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit: // select loop 탈출
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	quit <- true	// 종료 채널
	wg.Wait()
}
