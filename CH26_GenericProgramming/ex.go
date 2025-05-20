package main

import "fmt"

func print[T any](a T) {
	fmt.Println(a)
}

func main() {
	var a int = 10
	print(a)
	var b float32 = 3.14
	print(b)
	var c string = "hello"
	print(c)
}
