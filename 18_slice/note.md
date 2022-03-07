# 슬라이스

배열 메모리의 동적할당을 가능하게 해준다.

## 슬라이스

```go
var array [10]int
var slice []int
```

slice 사용을 위해서는 반드시 **초기화**가 필요

```go
package main

import "fmt"

func main() {
    var slice []int
    if (len(slice) == 0) {
        fmt.Println("slice is empty")
    }

    slice[1] = 10 // Panic!!
}
```

### `{}`를 이용한 슬라이스 초기화

```go
var slice1 = []int{1,2,3}
var slice2 = []int{1, 2:3, 4:5} // 1, 0, 0, 2, 0, 4
```

#### 다음에 주의!!

```go
var array = [...]int{1,2,3} // 길이가 3인 배열
var slice = []int{1,2,3} // 길이가 3인 슬라이스
```

### `make()`를 이용한 슬라이스 초기화

```go
var slice = make([]int, 3) // 길이가 3인 슬라이스 0,0,0
```

### 슬라이스 요소 접근

```go
var slice = make([]int, 3)
slice[1] = 5
```

### 슬라이스 순회

```go
var slice = []int{1,2,3}
// 각 요소에 10 더하기
for i := 0; i < len(silce); i++ {
    slice[i] += 10
}
// 각 요소에 2 곱하기
for i,v := range(slice) {
    slice[i] = v * 2
}
```

### 슬라이스 추가 : `append`

```go
var slice = []int{1,2,3}
slice := append(slice,4) // 1,2,3,4
// 여러값 추가
slice := append(slice,5,6,7) // 1,2,3,4,5,6,7
```

## 슬라이스 동작원리

```go
type SliceHader struct {
    Data uintptr    // 실제 배열을 가리키는 포인터
    Len int         // 요소 개수
    Cap int         // 실제 배열 길이
}
```

### 슬라이스와 배열의 동작 차이

- Golang의 모든 대입 값은 **복사**에서 이루어진다. 즉,
  - `changeArray`에서 `array2`는 array의 값을 복사한 배열이다.
  - 하지만, `changeSlice`에서의 `slice2`는 slice 구조상 포인터의 메모리 주소가 복사되기 때문에 똑같은 배열을 가리키게 된다.

```go
package main

import "fmt"

func changeArray(array2 [5]int) {
    array2[2] = 200
}

func changeSlice(slice2 []int) {
    slice2[2] = 200
}

func main() {
    array := [5]int{1,2,3,4,5}
    slice := []int{1,2,3,4,5}

    changeArray(array) // 1,2,3,4,5
    changeSlice(slice) // 1,2,200,4,5
}
```

#### `append()`를 사용할 때 발생하는 예기치 못한 문제 1

- `append` 함수가 호출되면 빈 공간이 있는지 확인하는데 이는, `남은 빈공간 = cap - len` 으로 계산됨

```go
func main() {
    slice1 := make([]int, 3, 5) // 0,0,0,0,0 len:3, cap:5
    slice1[0] = 1
    slice1[1] = 2
    slice1[2] = 3 // 1,2,3,0,0

    slice2 := append(slice1, 4,5) // 1,2,3,4,5 len:5, cap:5 (empty space = cap - len = 2)

    // slice1과 slice2의 구조체의 `len`값 차이가 발생한다.
    slice[1] = 100
    // - slice1: 1,100,3,4,5 {len:3, cap:5}
    // - slice2: 1,100,3,4,5 {len:5, cap:5}

    // 각 slice 빈공간의 인지 차이에 따라 의도한 값과 다르게 나온다.
    slice[1] := append(slice1, 500)

    // - slice1: 1,100,3,500 {len: 4, cap:5}
    // - slice2: 1,100,3,500,5 {len:5, cap:5}
}
```

#### `append()`를 사용할 때 발생하는 예기치 못한 문제 2

빈 공간이 없을 때 값을 추가하면 어떻게 될까?

- 슬라이스에 `append()`로 값을 추가시, `slice2`는 새로운 길이로 배열을 만들고 기존 `slice1`의 모든 값을 복사하고 맨 뒤에 4,5를 추가해서 `len:5, cap:6`의 슬라이스를 반환

```go
package main

import "fmt"

func main() {
    slice1 := []int{1,2,3}       // slice1: {len:3, cap:3}
    slice2 := append(slice1,4,5) // slice2: {len:5, cap:6}

    // 중간값 변경
    slice1[1] = 100
    // slice1: [1,100,3]
    // slice2: [1,2,3,4,5] <- 변경없음(다른 배열)
}
```

### 슬라이싱

- **슬라이싱** : 배열의 일부를 집어내는 기능
- return값 = 슬라이스

```go
array[startIndex:endIndex] // startIndex ~ endIndex-1까지의 슬라이스 반환

arr := [7]int{1,2,3,4,5,6,7}
arr[1:5] // [2,3,4,5] <- 배열의 일부를 가리키는 슬라이스를 반환(= 새로운 배열 생성 x)(=배열의 일부를 가리키는 포인터)
```

- [슬라이싱으로 배열 일부를 가리키는 슬라이스 만들기](./slicing/main.go)

#### 슬라이스를 슬라이싱

```go
slice1 := []int{1,2,3,4,5}
slice2 := slice1[1:2] // [2] len:1, cap:4

// 처음부터 슬라이싱
slice1 := []int{1,2,3,4,5}
slice2 := slice1[0:3] // [1,2] len:2, cap:5

// 끝까지 슬라이싱
slice1 := []int{1,2,3,4,5}
slice2 := slice1[2:len(slice1)] // [3,4,5] len:3, cap3

// 전체 슬라이싱 - 배열전체를 슬라이스로 만들경우 사용
slice1 := []int{1,2,3,4,5}
slice2 := slice1[:] // [1,2,3,4,5] len:5, cap:5

// 인덱스 3개로 슬라이싱해 cap 크기 조절
// slice[시작인덱스 : 끝인덱스 : 최대인덱스]
slice1 := []int{1,2,3,4,5}
slice2 := slice1[1:3:4] // [2,3] len:2, cap:3
```

### 유능한 슬라이싱 기능 활용

- [슬라이스 복제](./duplicate-slice/main.go)
- [요소 삭제](./delete-element/main.go)
- [요소 추가](./add-element/main.go)
- [슬라이스 정렬](./sort-slice/main.go)
