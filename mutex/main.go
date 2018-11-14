package main

import (
	"fmt"

	"github.com/SrcHndWng/go-learning-samples/mutex/models"
)

func main() {
	fmt.Println("Mutex Sample Start.")

	done := make(chan bool)
	account := models.NewAccount(1000)
	go account.Withdraw(700, done)
	go account.Deposite(500, done)
	<-done
	<-done

	fmt.Printf("Final balance = %d\n", account.Balance())
}
