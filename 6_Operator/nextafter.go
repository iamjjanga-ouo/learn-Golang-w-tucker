package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b // 값 비교
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f == %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.000000000000004
	b = 0.000000000000002
	c = 0.000000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
}

/*
0.100000000000000006 + 0.200000000000000011 == 0.300000000000000044
0.299999999999999989 == 0.300000000000000044 : true
7e-15 == 6.0000000000000005e-15 : false
*/
