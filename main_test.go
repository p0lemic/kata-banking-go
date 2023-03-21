package main

import (
	"testing"

	"github.com/p0lemic/banking-go/account"
)

func TestAccount_Deposit(t *testing.T) {
	account, err := account.NewAccount()

	if err != nil {
		t.Error(err)
	}

	account.Deposit(500)
	if account.Balance() != 500 {
		t.Errorf("Expected balance of 500, got %d", account.Balance())
	}
}

func TestAccount_Withdraw(t *testing.T) {
	account, err := account.NewAccount()

	if err != nil {
		t.Error(err)
	}

	account.Withdraw(500)
	if account.Balance() != -500 {
		t.Errorf("Expected balance of -500, got %d", account.Balance())
	}
}

func TestAccount_Balance(t *testing.T) {
	account, err := account.NewAccount()

	if err != nil {
		t.Error(err)
	}

	account.Deposit(500)
	account.Withdraw(100)
	if account.Balance() != 400 {
		t.Errorf("Expected balance of 400, got %d", account.Balance())
	}
}
