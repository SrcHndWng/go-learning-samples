package models

import (
	"fmt"
	"sync"
)

type Account struct {
	mutex   sync.Mutex
	balance int
}

func NewAccount(init int) *Account {
	fmt.Printf("init balance = %d\n", init)
	return &Account{balance: init}
}

func (a *Account) Deposite(value int, done chan bool) {
	a.mutex.Lock()
	fmt.Printf("balance = %d, deposite %d to account\n", a.balance, value)
	a.balance += value
	a.mutex.Unlock()
	done <- true
}

func (a *Account) Withdraw(value int, done chan bool) {
	a.mutex.Lock()
	fmt.Printf("balance = %d, withdraw %d from account\n", a.balance, value)
	a.balance -= value
	a.mutex.Unlock()
	done <- true
}

func (a *Account) Balance() int {
	return a.balance
}
