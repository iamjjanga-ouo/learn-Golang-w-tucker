// 숫자맞추기게임
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func getRandNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) // 0 ~ 99
}

var stdin = bufio.NewReader(os.Stdin)

func inputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n') // \n일때까지 읽음
	}
	return n, err
}

func main() {
	targetNumber := getRandNum()
	cnt := 1
	for {
		fmt.Printf("숫자값을 입력하세요.")
		n,err := inputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if n > targetNumber {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < targetNumber {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도 횟수:", cnt)
				break
			}
			cnt++
		}
	}
}