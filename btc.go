package main

// import "errors"

// type BTCAccount struct {
// 	balance int
// 	fee int
// }

// func NewBTCAccount() *BTCAccount {
// 	return &BTCAccount{
// 		balance: 0,
// 		fee: 300,
// 	}
// }

// func (nsb *BTCAccount) GetBalance() int {
// 	return nsb.balance
// }

// func (nsb *BTCAccount) Deposit(amount int) {
// 	nsb.balance += amount
// }

// func (nsb *BTCAccount) Withdraw(amount int) error {
// 	newBalance := nsb.balance - amount - nsb.fee
// 	if newBalance < 0 {
// 		return errors.New("insufficent funds")
// 	}
// 	nsb.balance = newBalance
// 	return nil
// }