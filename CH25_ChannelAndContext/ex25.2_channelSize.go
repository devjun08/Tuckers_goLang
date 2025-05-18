package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 9
	fmt.Println("Never print")
}
