package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

/*
- 정상 입력
1 2
2 1 2

- 비정상 입력
4 Hello
1 expected integer
*/
