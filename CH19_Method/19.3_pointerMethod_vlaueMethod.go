package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

// Point Method
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// Value Method
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// Return Original Value type method
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "JunYeop", "Kim"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance)
}
