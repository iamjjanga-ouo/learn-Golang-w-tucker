## 함수

```Go
func Add(a int, b int) {
  return a + b
}
```
- Go 함수의 특징
  - 첫글자가 대문자인 함수는 패키지를 외부로 공개

### 멀티 반환 함수
```Go
func Divide(a, b int) (int, bool) {
  if b == 0 {
    return 0, false
  } else {
    return a / b, true
  }
}

c, success := Divide(9,3) // 3 true
```

### 변수명 지정해서 반환
```Go
func Divide(a, b int) (result int, success bool) {
  if b == 0 {
    result = 0
    success = false
    return
  }
  result = a / b
  success = true
  return
}
```

### 재귀 호출
```Go
func printNo(n int) {
  if n == 0 {
    return
  }
  fmt.Println(n)
  printNo(n-1)
  fmt.Println("After", n)
}
func main() {
  printNo(3)
}

/*
3
2
1
After 1
After 2
After 3
*/
```