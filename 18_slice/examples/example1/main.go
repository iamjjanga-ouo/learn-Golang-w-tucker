package main

import "fmt"

func main() {
	array := [5]int{1,2,3,4,5}
	slice := array[1:3] // [2,3] {len:2, cap:4}

	slice = append(slice, 100) // [2,3,100] {len:3, cap:4}
	fmt.Println(array) // [1,2,3,100,5]
}