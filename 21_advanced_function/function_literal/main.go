package main

import (
	"fmt"
)

// 가독성 줄이기
type opFunc func(int, int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
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
