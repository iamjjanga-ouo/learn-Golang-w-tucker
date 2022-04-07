// 작업 시간을 설정한 컨텍스트
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 컨텍스트 생성, 5초후에는 context가 동작안함
	go PrintEverySecond(ctx)
	time.Sleep(10 * time.Second) // 10초뒤
	cancel()                     // 컨텍스트 캔슬

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
