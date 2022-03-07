package main

import "fmt"

func main() {
	array := [5]int{1,2,3,4,5}
	slice := array[1:2]
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice)) // len: 1, cap: 4 (cap = 배열의 총길이 - 처음index)

	array[1] = 100
	fmt.Printf("data: %v, len: %d, cap: %d\n", array, len(array), cap(array))
	fmt.Printf("data: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	// data: [1 100 3 4 5], len: 5, cap: 5
	// data: [100], len: 1, cap: 4

	slice = append(slice, 500)
	fmt.Printf("data: %v, len: %d, cap: %d\n", array, len(array), cap(array))
	fmt.Printf("data: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	// data: [1 100 500 4 5], len: 5, cap: 5
	// data: [100 500], len: 2, cap: 4
}