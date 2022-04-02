package main

import "fmt"

// 가독성 줄이기
type opFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) opFunc {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)
}
