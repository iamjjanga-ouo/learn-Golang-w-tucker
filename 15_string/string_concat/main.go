package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello"
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str));
	addr1 := stringHeader.Data

	str += "World"
	addr2 := stringHeader.Data

	str += "Welcome!"
	addr3 := stringHeader.Data

	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)

	// HelloWorldWelcome!
	// addr1:  49589e
	// addr2:  c000014090
	// addr3:  c0000160c0
}