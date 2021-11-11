## 구조체
- 형식
  ```Go
  type 타입명 struct {
    필드명 타입
    ...
    필드명 타입
  }
  ```

- 예시
  ```go
  type Student struct {
    Name  string
    Class int
    No    int
    Score float64
  }

  var t Student

  t.Name = "Tom"
  t.Class = "3"
  t.No = "21"
  T.Score = "16.04"
  ```

### 구조체 변수 초기화
- 초깃값 생략
  ```Go
  var t Student
  // Name: ""
  // Class: 0
  // No: 0
  // Score: 0.0
  ```
- 모든 필드 초기화
  ```go
  var t Student = Student{"Tom", "3", "21", "16.04"}

  // 여러줄
  var t Student = Student{
    "Tom",
    "3",
    "21",
    "16.04", // 쉼표 주의!
  }
  ```
- 일부 필드 초기화
  ```go
  var t Student = Student{Name:"Tom", No:"21"}
  ```

### 구조체를 포함하는 구조체
- 구조체의 필드로 다른 구조체 포함가능
  - [code](./struct_in_struct.go)
- 포함된 필드 접근방식
  - 포함되는 구조체의 필드에 접근하려면 `vipUser.UserInfo.Name` 이런 형식으로 접근해야한다. 하지만, 구조체에 다른 구조체 포함시 필드명을 생략하면 `.`으로 한번에 접근가능
  - [code](./access_include_field.go)
- 필드 중복시
  - `.`으로 정확히 몇 depth에 있는지 구분
  - [code](./overlap_field.go)


### 구조체 크기
```go
type User struct {
  Age   int
  Score float64
}
```
- Age(int, 8 바이트) + Score(float64, 8바이트) = 16바이트

### 구조체 복사
- Go 내부에서 필드 각각이 아닌 구조체 전체를 한 번에 복사
- [code](./copy_structure.go)

### 필드 배치 순서에 따른 구조체 크기 변화
- 작은 바이트의 필드를 앞에 배치
  ```go
  type User struct {
    Age   int32   // 4바이트
    Score float64 // 8바이트
  }

  user1 := User{23, 77.2}
  fmt.Println(unsafe.Sizeof(user)) // 16바이트
  ```
- 메모리 정렬을 고려한 배치
  ```go
  type User struct {
    Score float64 // 8바이트
    Age   int32   // 4바이트
  }

  user1 := User{23, 77.2}
  fmt.Println(unsafe.Sizeof(user)) // 12바이트
  ```
  - 메모리 정렬?
    - 컴퓨터가 데이터를 효율적으로 접근하고자 메모리를 일정 크기 간격으로 정렬
    - 64비트 컴퓨터에서 `int64` 데이터의 시작 주소가 100번지일 경우 100은 8의 배수가 아니기 때문에 104번지가 시작주소가 된다.
    - 위의 배치 풀이
      - 4바이트 + (실제 배치되지 않은 4바이트) + 8바이트 = 16바이트
      - 8바이트 + 4바이트 = 12바이트
- 메모리 낭비를 줄이는 방법?
  - 8바이트 보다 작은 필드는 8바이트 크기(단위)를 고려해서 몰아서 배치