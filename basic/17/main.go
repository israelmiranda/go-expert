package main

import "fmt"

type Account struct {
	Balance int
}

func NewAccount() *Account {
	return &Account{Balance: 0}
}

func (ac *Account) deposit(value int) {
	ac.Balance += value
}

func (ac *Account) withdrall(value int) {
	ac.Balance -= value
}

func main() {
	a := NewAccount()
	a.deposit(100)
	a.withdrall(100)

	fmt.Printf("Account balance: %d\n", a.Balance)
}
