package main

import "fmt"

func main() {
	// Integer Overflow
	var x int8 = 127 // 8비트 부호가 있는 정수 최대값
	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t= %4d, %08b\n", x, x)
	fmt.Printf("x + 1\t= %4d, %08b\n", x+1, x+1)
	fmt.Printf("x + 2\t= %4d, %08b\n", x+2, x+2)
	fmt.Printf("x + 3\t= %4d, %08b\n", x+3, x+3)

	// Integer Underfolw
	var y int8 = -128 // 8비트 부호가 있는 정수 최솟값
	fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1)
	fmt.Printf("y\t= %4d, %08b\n", y, y)
	fmt.Printf("y - 1\t= %4d, %08b\n", y-1, y-1)
}

/*
127 < 127 + 1: false
x       =  127, 01111111
x + 1   = -128, -10000000
x + 2   = -127, -1111111
x + 3   = -126, -1111110
-128 > -128 - 1: false
y       = -128, -10000000
y - 1   =  127, 01111111
*/
