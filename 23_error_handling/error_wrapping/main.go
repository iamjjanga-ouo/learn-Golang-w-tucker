package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
에러를 감싸서 새로운 에러를 만드는 경우
- 파일 읽기에서도 발생하는 에러도 필요
- 텍스트의 몇 번째 줄의 몇 번째 칸에서 에러가 발생하는지 알면 더유용
-> 파일 읽기에서 발생한 에러를 감싸고 그 바깥에 줄과 칸 정보를 넣기
*/

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) // 스캐너 생성
	scanner.Split(bufio.ScanWords)                      // 한 단어씩 끊어서 읽기
	// scanner.Split(bufio.ScanLines)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d, err:%w", pos, err)
	}

	pos = n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d, err:%w", pos, err)
	}
	return a * b, nil

}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("failed to scan")
	}

	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s, err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}
func main() {
	readEq("123 3")
	readEq("123 abc")
}
