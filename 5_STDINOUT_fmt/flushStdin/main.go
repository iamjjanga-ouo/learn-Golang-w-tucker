package main

import (
	"bufio" // io 담당 패키지
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 표준 입력 스트림 비우기
	} else {
		fmt.Println(n, a, b)
	}
	n, err = fmt.Scanln(&a, &b) // 다시 입력받기
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}

/*
Hello 4
expected integer
3 4
2 3 4
*/
