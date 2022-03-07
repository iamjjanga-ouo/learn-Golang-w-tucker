package main

import "fmt"

func main() {
	slice := []int{1,2,3,4,5,6}

	// 맨 뒤의 요소를 추가
	slice = append(slice,0)

	idx := 2

	// 맨 뒤부터 추가하려는 위치까지 값 하나씩 올리기
	for i := len(slice)-2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}

	// 값 변경
	slice[idx] = 100

	fmt.Println(slice)

	/* append() 함수로 코드 개선 */
	slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)

	/* 불필요한 메모리 사용이 없도록 코드 개선하기*/
	// 이전 append() 함수로 개선된 코드에서 [100,3,4,5,6]의 임시 메모리 공간을 사용

	slice = append(slice,0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 100
}