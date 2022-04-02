package main

import "fmt"

func Print(args ...interface{}) {
	for _, arg := range args {
		switch f := arg.(type) {
		case bool:
			val := f
			fmt.Println(val)
		case float64:
			val := f
			fmt.Println(val)
		case int:
			val := f
			fmt.Println(val)
		}
	}
}

func main() {
	Print(2, "Hello", 3.14)
}
