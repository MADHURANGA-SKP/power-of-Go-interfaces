package main

// import "errors"

// type NSBAccount struct {
// 	balance int
// }

// func NewNSBAccount() *NSBAccount {
// 	return &NSBAccount{
// 		balance: 0,
// 	}
// }

// func (nsb *NSBAccount) GetBalance() int {
// 	return nsb.balance
// }

// func (nsb *NSBAccount) Deposit(amount int) {
// 	nsb.balance += amount
// }

// func (nsb *NSBAccount) Withdraw(amount int) error {
// 	newBalance := nsb.balance - amount
// 	if newBalance < 0 {
// 		return errors.New("insufficent funds")
// 	}
// 	nsb.balance = newBalance
// 	return nil
// }