package main

import (
	"errors"
	"fmt"
)

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}
//----------------------------------------------------------------------
type BTCAccount struct {
	balance int
	fee int
}

func NewBTCAccount() *BTCAccount {
	return &BTCAccount{
		balance: 0,
		fee: 300,
	}
}

func (nsb *BTCAccount) GetBalance() int {
	return nsb.balance
}

func (nsb *BTCAccount) Deposit(amount int) {
	nsb.balance += amount
}

func (nsb *BTCAccount) Withdraw(amount int) error {
	newBalance := nsb.balance - amount - nsb.fee
	if newBalance < 0 {
		return errors.New("insufficent funds")
	}
	nsb.balance = newBalance
	return nil
}
//-------------------------------------------------------------------------------
type NSBAccount struct {
	balance int
}

func NewNSBAccount() *NSBAccount {
	return &NSBAccount{
		balance: 0,
	}
}

func (nsb *NSBAccount) GetBalance() int {
	return nsb.balance
}

func (nsb *NSBAccount) Deposit(amount int) {
	nsb.balance += amount
}

func (nsb *NSBAccount) Withdraw(amount int) error {
	newBalance := nsb.balance - amount
	if newBalance < 0 {
		return errors.New("insufficent funds")
	}
	nsb.balance = newBalance
	return nil
}
//---------------------------------------------------------------------------
func main() {
	myAccounts := []IBankAccount{
		NewBTCAccount(),
		NewNSBAccount(),
	}

	for _, account := range myAccounts{
		 account.Deposit(500)
		 if err := account.Withdraw(70); err != nil {
			fmt.Printf("ERR: %d\n", err)
		}

		balance := account.GetBalance()
		fmt.Printf("balance = %d\n", balance)
	}
}