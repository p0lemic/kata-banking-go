package main

import (
	"fmt"

	"github.com/p0lemic/banking-go/account"
)

func main() {
	account, err := account.NewAccount()

	if err != nil {
		panic(err)
	}

	account.Deposit(500)
	account.Withdraw(100)

	fmt.Println(account.PrintStatement())
}
