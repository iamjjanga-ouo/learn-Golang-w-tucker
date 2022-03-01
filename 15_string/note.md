## 문자열

---

### 큰따옴표와 백쿼트

```go
// 큰따옴표로 묶으면 특수 문자가 사용가능
str1 := "Hello\tWorld\t"
// 백쿼트로 묶으면 특수 문자 동작 x
str2 := `Hello\tWorld\t`

fmt.Println(str1) // Hello  World
fmt.Println(str2) // Hello\tWorld\t
```

큰따옴표와 백쿼트를 사용한 여러 줄 문자열 출력

```go
poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를,\n잎새에 이는 바람에도\n나는 괴로워했다."
poet2 := `죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를,
잎새에 이는 바람에도
나는 괴로워했다.`
```

<br>

### UTF-8

- Golang의 창시자 롭 파이크, 켄 톰슨이 고안한 문자코드
- UTF-16에서의 한문자에 2byte가 고정인것에 반해, (영문자+숫자+일부특수문자)를 1byte 그외에는 2~3byte로 표현하여 크기 절약 및 ASCII와 1:1대응
- Golang에서는 기본적으로 한글이 지원된다.

<br>

### rune 타입으로 한 문자 담기

- UTF-8은 1~3byte 크기, 하지만 Golang에서는 3btye단위가 없기때문에 4byte 정수 타입인 int32 타입의 별칭으로 `rune`을 사용

  ```go
  type rune int32

  var char rune = '한'

  fmt.Printf("%T\n", char)    // int32
  fmt.Println(char)           // 54620
  fmt.Printf("%c\n", char)    // 한
  ```

<br>

### `len()` : 문자열 크기

한글은 글자당 3byte의 크기를 가진다.

```go
str1 := "가나다라마"
str2 := "abcde"

fmt.Printf("%d\n", len(str1)) // 15
fmt.Printf("%d\n", len(str2)) // 5
```

<br>

### `[]rune` 타입변환으로 글자 수

- string 타입과 `[]rune` 타입은 모두 문자들의 집합을 나타내기 때문에 상호 타입 변환가능

```go
func main() {
	str := "Hello 월드"
	runes := []rune(str)

	fmt.Printf("len(str) = %d\n", len(str)) 		// len(str) = 12
	fmt.Printf("len(runes) = %d\n", len(runes))		// len(runes) = 8
}
```

<br>

## 문자열순회

---

### 1. 인덱스를 사용한 바이트 단위 순회

```go
func main() {
    str := "Hello 월드!"

    for i := 0; i < len(str); i++ {
        fmt.Printf("타입: %T / 값: %d / 문자값: %c\n", str[i], str[i], str[i])
    }
}

// 타입: uint8 / 값: 72 / 문자값: H
// ...
// 타입: uint8 / 값: 236 / 문자값: ì  <- 3바이트 한글은 깨짐
```

바이트로 순회시 한글이 깨지는 이슈를 다음 2가지 방법으로 대체 가능하다.

1. `[]rune` 타입변환 후 한글자씩 순회
2. `range` 키워드를 이용해 한글자씩 순회

<br>

### 2. `[]rune` 타입변환 후 한 글자씩 순회

```go
str := "Hello 월드!"
runes := []rune(str)

for i := 0; i < len(runes); i++ {
    fmt.Printf("타입: %T / 값: %d / 문자값: %c\n", runes[i], runes[i], runes[i])
}
// 타입: int32 / 값: 72 / 문자값: H
// ...
// 타입: int32 / 값: 50900 / 문자값: 월
// 타입: int32 / 값: 46300 / 문자값: 드
```

<br>

### 3. `range` 키워드를 이용해 한 글자씩 순회

```go
func main() {
    str := "Hello 월드!"

    for _,v := range(str) {
        fmt.Printf("타입: %T / 값: %d / 문자값: %c\n", v, v, v)
    }
}
```

## 문자열 합치기

---

```go
func main() {
    str1 := "Hello"
    str2 := "World"

    str3 := str1 + " " + str2
    fmt.Println(str3)

    str1 += " " + str2
    fmtPrintln(str1)
}
```

## 문자열 비교

---

- `==`, `!=`를 사용해 비교

<br>

## 문자열 대소비교

---

- `>`, `<`, `<=`, `>=`를 통해 대소 비교 (UTF-8의 값이 큰지, 작은지 비교)

<br>

## 문자열 구조

---

```go
type StringHeader struct {
    Data uintptr    // 문자열 데이터가 있는 메모리주소의 포인터
    Len int         // 문자열의 길이
}
```

<br>

### string끼리 대입

```go
package main

import "fmt"

func main() {
    str1 := "안녕하세요. 한글 문자열입니다."
    str2 := str1

    // str2가 str1의 구조체 값이 복사가 된다.
    // 각 필드의 `Data` 포인터값, `Len` 값이 복사 된다.
    // 문자열 자체의 복사 x

    stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1)) // Data값 추출
    stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2)) // Data값 추출

    fmt.Println(stringHeader1) // &{4983854 12}
    fmt.Println(stringHeader2) // &{4983854 12}
}
```

<br>

## 문자열은 불변이다.

---

- 불변 = string 타입이 가리키는 문자열 일부만 변경 x

```go
var str string = "Hello World"
str = "How are you?" // OK
str[2] = 'a' // Error
```

다른 방법으로 변환은 가능

```go
package main

import "fmt"

func main() {
    var str string = "Hello World"
    var slice []byte = []byte(str)

    slice[2] = 'e'

    fmt.Println(str)    // Hello World
    fmt.Println(slice)  // Heelo World

    stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
    sliceHeader := (*reflect.StringHeader)(unsafe.Pointer(&slice))

    fmt.Printf("str : \t%x\n", stringHeader.Data)   // str: 4beb6f
    fmt.Printf("slice : \t%x\n", slcieHeader.Data)  // slice:   c00006af10
}
```

<br>

### 문자열 합산에서 보여지는 문자열의 변화

- `string` 타입간의 합 연산에서 문자열이 어떻게 변화되는지 확인
- [Code](./string_concat/main.go)

<br>

### Builder : string 합연산 패키지

- string 합 연산을 빈번하게 사용하면 메모리의 낭비가 잇음
- `Builder`을 통해 메모리 낭비를 줄일 수 있다.
  - `strings.Builder`의 `WriteRune()` 메서드는 문자를 더할때 매번 메모리르 새로 생성하지 않고 기존 메모리 공간에 빈자리가 있으면 더하게 된다.
- [Code](./builder_example/main.go)

<br>

### 왜 문자열은 불변 원칙을 지키려할까?

- 빈번한 연산 시 메모리 낭비되는 것임에도 불구하고 문자열 불변 원칙을 지키려고 할까?
  -> 예기치 못한 버그를 방지
- 만약 불변의 원칙이 없다면 문자열이 언제라도 변화할 수 있어서 `string` 타입 변수를 안심하고 사용할 수 없음
