## IF문

### 쇼트 서킷
- &&(AND) 조건에서는 좌변이 false면 우변을 검사하지 않고 false 판단
- ||(OR) 조건에서는 좌변이 true면 우변을 검사하지 않고 true 판단

### if 초기문; 조건문
```Go
if filename, success := UploadFile(); success {
  fmt.Println("Upload Success", filename)
} else {
  fmt.Println("Upload Failed")
}
```
- 초기문에서의 변수 범위는 if 문 `안`으로 한정됨.

## Switch문
- 다른 언어랑 비슷

### 한번에 여러값 비교
```Go
switch day {
  case "monday", "tuesday":
    fmt.Println("월,화는 수업 가는 날입니다.")
  case "wednesday", "thursday", "friday":
    fmt.Println("수,목,금은 실습 가는 날입니다.")
}
```

### 조건문 비교
- Switch문의 동작응 if문 처럼 true가 되는 조건문으로 검사할 수 있음
- [예제 코드](./switch_like_if.go)

### switch 초기문
```Go
switch age := GetMyAge(); age {
case 10:
  ...
case 33:
  ...
default:
  ...
}
```
- 비교값이 `true`일 경우 생략가능
  ```Go
  switch age := GetMyAge(); {
    ...
  }
  ```

### break와 fallthrough 키워드
- 다른 언어와 다르게 `break`를 사용하지 않더라도 Go언어에서는 `case` 하나 실행 후 switch문을 자동으로 빠져나올 수 있음.
- 만약 case문 실행 후 빠져나오지 않고 다음 case문까지 실행하고 싶은경우 -> `fallthrough` 키워드를 사용
- 하지만, 코드를 보는부분에 좋지않아서 되도록 사용하지 않는 것을 권장
  ```Go
  a := 3

  switch a {
  case 1:
    ...
    break
  case 2:
    ...
  case 3:
    ...
    fallthrough // !
  case 4:
    ...
  default:
  }
  ```