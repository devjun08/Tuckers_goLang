package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello"
	str2 := str1

	// unsafe 포인터 타입으로 변환 후 reflect 안에 있는 stringHeader 구조체 포인터 타입으로 변환
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}
