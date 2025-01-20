package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	userPointer := NewUser("AAA", 23)

	fmt.Println(userPointer)
}

// 탈출 분석(escape analysis)으로 이 코드는 실행이 됨. Golang의 특징
