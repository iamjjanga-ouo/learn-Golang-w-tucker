/*
1. handler function
2. query string
*/
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query() // 쿼리 인수 가져오기
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

var port int

func main() {
	port = 3000
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	http.HandleFunc("/bar", barHandler)

	fmt.Printf("Server is listening ... %d port", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
