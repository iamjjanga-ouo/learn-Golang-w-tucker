# 패키지

## 유용한 패키지

- [gopkg](https://pkg.go.dev/github.com/astaxie/gopkg)
- [awesome-go](https://github.com/avelino/awesome-go)

## 패키지 사용

```go
import "fmt"

// multiple packages
import (
    "fmt"
    "strings"
)
```

- **외부 노출 여부**는 첫글자를 대문자 사용여부에 따라 달라짐

```go
package main

import (
    "fmt"
    "text/template"
    htemplate "html/template" // 겹치는 이름은 별칭으로 사용
    "database/sql"
    _ "github.com/mattn/go-sqllite3" // 사용하지 않는 패키지 포함 (직접사용하지는 않고, database/sql에 부가적으로 사용)
)

func main() {
    template.New("foo")
    htemplate.New("foo")
}
```

## Go Module

- Go 1.16부터 Go 모듈사용이 **기본값**이 됨
- 이전까지의 Go Project는 `$GOPATH/src` 디렉터리 하위에서 모듈을 관리했음
- Go Module이 기본이 되면서 Go Module 디렉터리 하위에는 `go.mod`(Module 이름 및 필요한 외부 패키지 정보) 파일과 `go.sum`(외부 저장소 패키지 버전 정보)를 가지고 있어야한다.
  ```go
  go mod init [패키지명]
  ```
- [code](./usepkg/usepkg.go)

## 패키지명 & 패키지 외부 공개
- 쉽고 간단하게 패키지 이름 작성
- 모든 문자를 **소문자**
- 첫글자가 **대문자** 인 변수, 메서드는 외부공개

## 패키지 초기화
1. package를 import하는 순간 compiler는 package 내 전역변수를 초기화
2. 패키지에 `init()` 함수가 있다면 호출해 패키지를 초기화. (`init()` 함수는 반드시 **input argument**가 없고 **return** 값도 없는 함수여야함.)
3. `init()` 함수 기능만 사용하기를 원할 경우 `_`을 이용해서 import
    ```go
    _ "github.com/mattn/go-sqlite3"
    ```
