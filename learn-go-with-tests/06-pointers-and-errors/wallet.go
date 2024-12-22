package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("can't withdraw, insufficient funds.")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amt Bitcoin) {
	// fmt.Printf("address of bal in Deposit is %p\n", &w.balance)
	w.balance += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
