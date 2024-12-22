package wallet

import "fmt"

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

func (w *Wallet) Withdraw(amt Bitcoin) {
	w.balance -= amt
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
