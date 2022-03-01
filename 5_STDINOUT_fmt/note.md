## 5 : fmt 패키지를 이용한 표준 입출력

### fmt

- STDOUT

  전반적으로 C언어 공부를 비슷해서 스킵

  ```Go
  Print()
  Println()
  Printf('%d %s %v', value...) // format (%v는 value에 맞춰 기본 타입으로 해줌.)
  ```

- STDIN

  - Scan()
    ```Go
    func Scan(a ...interface{}) (n int, err error)
    ```
    - Return
      - 성공 : 입력한 값의 개수
      - 실패 : 에러
    - [Scan문 예제](scan/main.go)
  - Scanf()
    ```Go
    func Scanf(format string, a ...interface{}) (n int, err error)
    ```
    - [Scanf문 예제](scanf/main.go)
  - Scanln()
    - 한 줄로 입력받고 공백으로 나눔.

- 키보드 입력과 Scan() 함수의 동작원리
  - FIFO 입력 데이터 처리 순서
    ```
    Keyboard - [\n,4, ,o,l,l,e,H ] -> Computer
    ```
    - 첫글자 'H'를 읽어서 숫자가 아니면 에러를 만들면서 입력 스트림을 비움
    - [입력스트림 비우는 예시](flushStdin/main.go)
