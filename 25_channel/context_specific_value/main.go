// 특정 값을 설정한 컨텍스트
package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)
	// 컨텍스트에 값을 추가한다.
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int) // Value()의 반환타입은 빈 인터페이스라서 (int)타입으로 변환하여 사용
		fmt.Printf("Square:%d", n)
	}
	wg.Done()
}

// 취소도 되면서 값도 설정하는 컨텍스트는?
/*
ctx, cancel := context.WitCancel(context.Background())	// 취소 기능이 있는 컨텍스트 생성
ctx = context.WithValue(ctx, "number", 9)				// 값을 설정한 컨텍스트를 만듦
ctx = context.WithValue(ctx, "keyword". "grey")			// 컨텍스트를 여러번 감싸면 여러 값을 설정할 수 있음
*/
