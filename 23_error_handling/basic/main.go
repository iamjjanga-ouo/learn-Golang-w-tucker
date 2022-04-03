package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) // 파일 읽기
	if err != nil {
		return "", err // 에러 나면 에러 반환
	}
	defer file.Close() // 함수 종료 직전 파일 닫기

	rd := bufio.NewReader(file) // 파일 내용 읽기
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) // 파일 생성
	if err != nil {                  // 에러 나면 에러 반환
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, line) //파일에 문자열 쓰기
	return err
}

const wd string = "/workspace/Gode-server/learn-Golang-w-tucker/23_error_handling/basic"
const filename string = "data.txt"

func main() {
	path := filepath.Join(wd, filename)
	line, err := ReadFile(path) // 파일 읽기 시도
	if err != nil {
		err = WriteFile(path, "This is writeFile") // ReadFile에러시 파일 생성 및 쓰기
		if err != nil {
			fmt.Println("파일 생성에 실패", err)
			return
		}
		line, err = ReadFile(path) // 생성한 파일 읽기
		if err != nil {
			fmt.Println("파일 읽기에 실패", err)
			return
		}
	}
	fmt.Println("파일 내용 : ", line) // 출력
}
