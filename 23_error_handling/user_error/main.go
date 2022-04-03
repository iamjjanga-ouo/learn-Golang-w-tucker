package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야합니다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

func main() {
	number := float64(-2)
	sqrt, err := Sqrt(number)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Sqrt(%f) = %v\n", number, sqrt)
}
