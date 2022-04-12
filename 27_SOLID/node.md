# 객체지향 설계 5가지 원칙 : SOLID
- 단일 책임 원칙 (Single responsibility principle, SRP)
- 개방-폐쇄 원칙 (Open-closed principle, OCP)
- 리스코프 치환 원칙 (Liskov subsitution principle, LSP)
- 인터페이스 분리 원칙 (Interface segregation principle, ISP)
- 의존 관계 역전 원칙 (Dependency inversion principle, DIP)

## 좋은설계 나쁜설계
- 나쁜 설계
    - `상호 결합도가 매우 높고 응집도가 낮다`
        - `상호 결합도가 높다` : 서로 강하게 결합되어 떼어낼 수 없다. 모듈의 수정이 다른 모듈로 전파되어 예기치 못한 문제가 생기고 코드 재사용성을 낮추게 됨.
        - `응집도가 낮다` : 하나의 모듈이 스스로 자립하지 못한다. 모듈이 스스로 완성되지 못하고 다른 모듈에 의존적인 관계
- 좋은 설계
    - `상호 결합도가 낮고 응집도가 높은 설계`

## 단일 책임 원칙
- "모든 객체는 책임을 하나만 져야 한다."

### 이점
- 코드의 재사용성

### 예시
```
<<interface>>
Report         <----- ReportSender
    ^
    |
FinanceReport
```

```go
type Report struct {
    Report() string
}

type FinanceReport struct {
    report string
}

func (r *FinanceReport) Report() string {
    return r.report
}

type ReportSender struct {
    ...
}

func (s *ReportSender) SendReport (report Report) {
    // Report 객체를 인수로 받음
    ...
}
```

## 개방-폐쇄 원칙
- "확장은 열려있고, 변경에는 닫혀있다."

### 이점
- 상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 된다.

### 예시
```go
// 개방-폐쇄가 지켜지지 않은 코드
func SendReport(r *Report, method SendType, receiver string) {
    switch method {
    case Email:
        // 이메일 전송
    case Fax:
        // 팩스 전송
    case PDF:
        // pdf 파일 전송
    case Printer:
        // 프린팅
    ...
    }
}
```

```go
// 개방-폐쇄가 지켜진 코드
type ReportSender interface {
    Send(r *Report)
}

type EmailSender struct {
}

func (e *EmailSender) Send(r *Report) {
    // 이메일 전송
}

type FaxSender struct {
}

func (f *FaxSender) Send(r *Report) {
    // 팩스 전송
}
```

## 리스코프 치환 원칙
- "q(x)를 타입 T의 객체 x에 대해 증명할 수 있는 속성이라 하자. 그렇다면 S가 T의 하위 타입이라면 q(y)는 타입 S의 객체 y에 대해 증명할 수 있어야한다."

### 이점
- 예상치 못한 작동을 예방할 수 있다.

### 예시
```go
type T interface {
    Something()
}

type S struct {
}

func (s *S) Something() { // T 인터페이스 구현
}

type U struct {
}

func (u *U) Something() { // T 인터페이스 구현
}

func q(t T){
}

var y = &S{}
var u = &U{}
q(y)
q(u)
```

## 인터페이스 분리 원칙
- "클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다."

### 이점
- 인터페이스를 분리하면 불필요한 메서들과 의존 관계가 끊어져 더 가볍게 인터페이스를 이용할 수 있다.

### 예시
- `Report` 인터페이스는 4개의 메서드를 포함
- `SendReport` 함수는 Report 인터페이스가 포함한 4개의 메서드중 `Report()`만 사용 즉, 인터페이스 이용자에게 불필요한 메서드를 포함
```go
type Report interface {
    Report() string
    Pages() int
    Author() string
    WrittenDate() time.Time
}

func SendReport(r Report) {
    send(r.Report())
}
```

- ISP에 입각
```go
type Report interface {
    Report() string
}

type WrittenInfo interface {
    Pages() int
    Author() string
    WritteDate() time.Time
}

func SendReport(r Report) {
    send(r.Report())
}
```

## 의존 관계 역전 원칙
- "상위 꼐층이 하위 계층에 의존하는 전통적인 의존 관계를 반전(역전)시킴으로써 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다."

### 원칙
1. "상위 모듈은 하위 모듈에 의존해서는 안 된다. 둘 다 추상 모듈에 의존해야한다."
2. "추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 모듈은 추상 모듈에 의존해야 한다."

### 이점
- 구체화된 모듈이 아닌 추상 모듈에 의존함으로써 확장성이 증가
- 상호 결합도가 낮아져서 다른 프로그램으로 이식성이 증가

### 원칙 1
Top-Down으로 사고하려는 경향
```
    전송
     |
  |-----|
키보드  네트워크
```

탑다운 방식이 아닌 새로운 방식으로 역전(invert) 해야함
```
입력 interface   <----전송----> 출력 interface
    ^                            ^
    |                            |
   키보드                         네트워크
```

### 원칙 2
ex) 메일이 수신되면 알람이 울린다고 가정
```go
type Mail struct {
    alarm Alarm
}

type Alarm struct {
}

func (m *Mail) OnRecv() { // OnRecv() 메서드는 메일 수신 시 호출 됨
    m.alarm.Alarm()       // 알림일 울립니다.
}

메일 *------ 알림 // 메일이 알림을 소유
```

의존 관계 역전 원칙에 입각해 바꿈
```
Event interface *------- EventListener interface
     ^                              ^
     |                              |
     |                              |
     메일                            알람
```

```go
type Event interface {
    Register(EventListener)
}

type EventListener interface {
    OnFire()
}

type Mail struct {
    listener EventListener
}

func (m *Mail) Register(listener EventListener) { // Event 인터페이스 구현
    m.listener = listener
}

func (m *Mail) OnRecv() { // 등록된 listener의 OnFire() 호출
    m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {  // EventListener 인터페이스 구현
    // 알람
    fmt.Println("알람! 알람!")
}

var mail = &Mail{}
var listener EventListener = &Alarm{}

mail.Register(listener)
mail.OnRecv()           // 알람이 울리게 됩니다.
```
