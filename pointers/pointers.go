package pointers

import (
	"errors"
	"fmt"
)

// ErrNoMoney money error message
var ErrNoMoney = errors.New("no money")

// Stringer strings things
type Stringer interface {
	String() string
}

// Bitcoin type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit money to wallet
func (w *Wallet) Deposit(coins Bitcoin) {
	w.balance += coins
}

// Withdraw money from wallet
func (w *Wallet) Withdraw(coins Bitcoin) error {
	if coins > w.balance {
		return ErrNoMoney
	}
	w.balance -= coins
	return nil
}

// Balance of wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
