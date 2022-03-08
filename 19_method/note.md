# 메서드

## 메서드 선언

```go
// func (Receiver) FunctionName ReturnType {}
func (r Rabbit) info() string {
    return "토끼입니다."
}
```

```go
package main

import "fmt"

type account struct {
    balance int
}

// 일반 함수 표현
func withdrawFunc(a *account, amount int) {
    a.balance -= amount
}

// 메서드 표현
func (a *account) withdrawMethod(amount int) {
    a.balance -= amount
}

func main() {
    a := &account{ 100 }

    withdrawFunc(a, 30)

	a.withdrawMethod(30)

	fmt.Printf("%d\n", a.balance)
}
```

#### Tips

메서드 정의는 같은 패키지 내 어디에도 위치할 수 있지만, 리시버 타입이 선언된 파일 안에 정의하는게 일반적인 규칙

- `type Student struct {}` 구조체를 student.go에 정의
- Student의 메서드들은 모두 student.go에 모아 정의

### 별칭 리시버 타입

모든 로컬 타입이 리시버 타입으로 가능하기 때문에, 별칭 타입도 리시버 OK / 메서드를 가질 수 있음

```go
package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
    return int(a) + b
}

func main() {
    var a myInt = 10
    fmt.Println(a.add(30)) // 40
    var b int = 20
    fmt.Println(myInt(b).add(50)) // 70
}
```

## 메서드는 왜 필요한가?

일반 함수 표현에서 argument를 앞으로 가져오는 (receiver)로 옮겼을 뿐 동일한 동작을 한다. 복작하게 Method를 만든이유는?

- 일반 함수 표현과 메서드 표현의 결정적 차이는 **"소속"**에 있다.
- 위의 예시에서 `withdrawFunc` 함수는 `*account` 타입과 분리되어 있지만, `withdrawMethod()` 메서드는 `*account` 타입에 속한 메서드 즉,`withdrawMethod()`는 account의 **기능**
- 좋은 프로그래밍에서의 결합도(Coupling, 객체간의 의존관계)은 낮추고, 응집도(cohesion, 모듈 내 요소들의 상호 관련성)은 높여야함.

### 객체지향: 절차 중심에서 관계 중심으로 변화

과거의 절차 중심 프로그래밍 (flowchart를 중시) -> 관계 중심 프로그래밍 (class diagram)으로의 변화가 있었음.

- Go 언어에서는 클래스와 상속을 지원하지 않고 **메서드**와 **인터페이스**만을 지원
- 일부는 OOP언어가 아니라고 말하지만, OOP 언어의 구분은 클래스와 상속 지원 유무가 아닌 **객체 간의 상화관계 중심으로 프로그래밍을 할 수 있는가? 없는가?**에 있음.

## 포인터 메서드 vs 값 타입 메서드

- [예제 코드](./pointer_n_value_method/main.go)

### 포인터 변수의 값 타입 메서드 호출

`mainA.withdrawValue(20)` `mainA`는 `*account` 포인터 변슁고 `withdrawValue`는 account 값 타입을 리시버로 받는 메서드 이때, `(*mainA).withdrawValue(20)`으로 호출하는게 맞지만 Go 언어에서는 자동으로 mainA 값으로 변환하여 호출함.

### 값 변수의 포인터 메서드 호출

`mainB.withdrawPointer(30)`에서 mainB는 account 값 타입 변수이고 `withdrawPointer()` 메서드는 `*account` 포인터를 리시버로 받는 메서드. 역시 값 타입인 `mainB`는 바로 호출할 수 없고, `(&mainB).withdrawPointer(30)`와 같이 주소 연산자를 사용해서 포인터로 변환 후에 호출해야합니다. 하지만 Go 언어에서는 자동으로 `mainB`의 메모리 주솟값으로 변환하여 호출

### 언제 포인터 메서드? 언제 값 타입 메서드?

| 포인터 메서드                    | 값 타입 메서드                                                                |
| -------------------------------- | ----------------------------------------------------------------------------- |
| 내부에서 리서버의 값을 변경 가능 | 호출하는 쪽과 메서드 내부의 별도 인스턴스로 독립(내부에서 리시버의 값 변경 x) |
| 인스턴스 중심                    | 값 중심                                                                       |
