/*
1. rsa:2048 방식의 키를 생성해 private key를 생성해 localhost.key로 저장하고 인증파일은 localhost.csr로 저장
```sh
openssl req -new -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr
```

2. 대표적인 인증 알고리즘 x509를 사용해 1년짜리 인증서(localhost.crt)발급
```sh
penssl x509 -req -days 365 -in localhost.csr -signkey localhost.key -out localhost.crt
```
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	err := http.ListenAndServeTLS(":3000", "localhost.crt", "localhost.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
