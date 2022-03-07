package main

import "fmt"

// POP을 구현한 예제
func main() {
	slice := []int{1,2,3,4,5,6}
	t, slice := slice[len(slice)-1], slice[:len(slice)-1]

	fmt.Println(t, slice) // 6 [1,2,3,4,5]
}