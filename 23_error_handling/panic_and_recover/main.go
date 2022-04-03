// panic은 호출 순서를 거슬러 올라가며 전파
// 함수 호출 과정 : main() -> f() -> g() -> h()
// h()함수에서 패닉이 발생하면 호출 순서는 h() -> g() -> f() -> main()
// recover() 함수가 호출하여 panic을 복구할 수 있음.

package main

import (
	"fmt"
	"net"
)

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨.")

	// recover() 결과
	// func recover() interface{}

	// 특정 타입에 대한 panic을 recover하는 경우
	if _, ok := recover().(net.Error); ok {
		fmt.Println("r is net.Error Type")
	}
}
