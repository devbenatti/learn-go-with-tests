package pointers

import (
	"errors"
	"fmt"
)

// ErrInsufficientBalance is a error message for insufficient balance
var ErrInsufficientBalance = errors.New("expect unable to withdraw: insufficient balance")

// Bitcoin represents a data structure of a bitcoin
type Bitcoin int

// Wallet represents a data structure of a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit makes a deposit in the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns a balance for wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw receive an amout and make a withdraw in a wallet
func (w *Wallet) Withdraw(amout Bitcoin) error {
	if amout > w.balance {
		return ErrInsufficientBalance
	}
	w.balance -= amout
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
