# 인터페이스

## 인터페이스 선언

```go
type DuckInterface interface {
    Fly()                   // 메서드
    Walk(distance int) int  // 메서드
}
```

### `{}`안에 메서드 사용시 유의사항

1. 메서드는 반드시 **메서드명**이 있어야함
2. 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없음 (Overloading x)
3. 인터페이스에서는 **메서드 구현을 포함하지 않음**

```go
package main

import "fmt"

type Stringer interface {
    String() string
}

type Student struct {
    Name    string
    Age     int
}

func (s Student) String() string {
    return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
    student := Student{ "철수", 12}
    var stringer Stringer

    stringer = student // stringer 값으로 student 대입

    fmt.Printf("%s\n", stringer.String()) // stringer의 String() 호출
}
```

- Go 언어에서는 `~er`을 붙여서 인터페이서 명을 만드는 것을 권장.

## 인터페이스 왜 쓰나?

- 인터페이스는 객체 지향 프로그래밍에 아주 중요한 역할
- 인터페이스를 이용하면 **구체화된 객체가 아닌 인터페이스만 가지고 메서드를 호출할 수 있음**.
- 큰 코드 수정없이 구체화된 객체를 바꿔서 사용할 수 있게됨.
- [코드 에제](./why_use_interface/main.go)

## 덕타이핑(duck typing)

- Go 언어에서 어떤 타입이 인터페이스를 포함하고 있는지 여부를 결정할 때 덕 타이핑 방식을 사용
- 덕 타이핑 : 타입 선언 시 인터페이스 구현 여부를 명시적으로 나타낼 필요 없이 인터페이스에 정의한 메서드 포함 여부만으로 결정하는 방식
- 사용자 중심의 프로그래밍을 가능하게한다.

```go
type Stringer interface {
    String() string
}

type Student struct {}

func (s *Student) String() string {} // 인터페이스에 정의한 메서드를 포함

// 덕타이핑이 없었더라면
type Student struct implements Stringer {}
```

## 인터페이스 기능 더 알기

- 포함된 인터페이스
- 빈 인터페이스
- 인터페이스 기본값

### 언터페이스를 포함하는 인터페이스

```go
type Reader interface {
    Read() (n int, err error)
    Close() error
}

type Writer interface {
    Write() (n int, err error)
    Close() error
}

// Reader, Writer 인터페이스의 메서드 집합을 모두 포함하는 ReadWriter 인터페이스
// Read(), Writer(), Close()를 가짐
type ReadWriter interface {
    Reader
    Writer
}
```

### 빈 인터페이스 `interface {}`를 인수로 받기

- 메서드를 가지고 있지 않은 빈 인터페이스
- 모든 타입이 빈 인터페이스로 쓰일 수 있음.
- 어떤 값이든 받을 수 있는 함수, 메서드, 변숫값을 만들 때 사용
- [코드 예제](./empty_interface/main.go)

### 인터페이스 기본값 nil

- 기본값 : 유효하지 않은 메모리 주소를 나타내는 `nil`

```go
package main

type Attacker interface {
    Attack()
}

func main() {
    var att Attacker
    att.Attack()        // panic: runtime error: invalid memory address or nil pointer dereference
}
```

## 인터페이스 변환하기

### 구체화된 다른 타입으로 타입 변환하기

```go
var a Interface
t := a.(ConcreteType) // 인터페이스 변수에 `.(타입)`을 붙여 변환
```

- [코드 예시](./convert_interface/main.go)

### 다른 인터페이스로 타입 변환하기

- 변경되는 인터페이스가 변경전 인터페이스를 포함하지 않아도 된다.
- 하지만, 인터페이스가 가리키고 있는 실제 인스턴스가 변환하고자 하는 다른 인터페이스를 포함해야한다.

```go
var a AInterface = ConcreteType{}
b := a.(BInterface)
```
