## 상수

- 상수 종류

  - 불리언
  - 룬(rune)
  - 정수
  - 실수
  - 복소수
  - 문자열

- 형식

  ```Go
  const ConstValue int = 10
  ```

- [상수가 좌변으로 값의 재할당이 일어났을 때 에러 예시](constant_error/main.go)

### iota로 간편하게 열거값 사용하기

- 상수를 코드의 값으로 사용할 때 그냥 1,2,3,.. 처럼 1씩 증가하도록 정의할 때 사용
  ```Go
  const (
    Red   int = iota    // 0
    Blue  int = iota    // 1
    Green int = iota    // 2
  )
  ```
- iota는 생략이 가능하다.
  ```Go
  const (
    C1 uint = iota + 1  // 0 + 1 = 1
    C2                  // 1 + 1 = 2
    C3                  // 2 + 1 = 3
  )
  ```

### 타입이 없는 상수

- 타입이 없는 상수를 선언할 수 있다.
- [예시 코드](notype_constant/main.go)
