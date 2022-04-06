// [sync]Sub goroutine이 종료되는 것을 기다리기
/*
var wg sync.WaitGroup

wg.Add(3)	// 작업 개수 설정
wg.Done()	// 작업이 완료될 때마다 호출
wg.Wait()	// 모든 작업이 완료될때까지 대기
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 10000)
	}

	wg.Wait()
	fmt.Println("모든 계산이 완료되었습니다.")
}
