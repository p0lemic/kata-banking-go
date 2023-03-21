package account

import (
	"fmt"
	"time"
)

type Account struct {
	balance    int
	operations []Operation
}

func NewAccount() (Account, error) {
	return Account{
		balance:    0,
		operations: make([]Operation, 0),
	}, nil
}

func (a *Account) Deposit(amount int) {
	a.balance = a.balance + amount

	a.operations = append(a.operations, Operation{
		date:    time.Now(),
		amount:  fmt.Sprintf("+%d", amount),
		balance: a.balance,
	})
}

func (a *Account) Withdraw(amount int) {
	a.balance = a.balance - amount

	a.operations = append(a.operations, Operation{
		date:    time.Now(),
		amount:  fmt.Sprintf("-%d", amount),
		balance: a.balance,
	})
}

func (a *Account) PrintStatement() string {
	output := ""

	for _, operation := range a.operations {
		line := fmt.Sprintln(operation.date.Format("2006-01-02"), operation.amount, operation.balance)
		output = output + line
	}

	return output
}

func (a *Account) Balance() int {
	return a.balance
}
