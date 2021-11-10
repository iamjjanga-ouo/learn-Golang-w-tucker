package main

import "fmt"

func main() {
	const C int = 10

	const b int = C * 20
	C = 20
	fmt.Println(&C)
}

// 8_Contants/contant_error.go:9:4: cannot assign to C (declared const)
