// Package account implements an API for manipulating bank accounts.
package account

import "sync"

// Account represents a bank account.
type Account struct {
	balance int64
	closed  bool
	mux     sync.Mutex
}

// Open opens a new Account and returns a reference to it.
func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}

	return &Account{balance: initalDeposit}
}

// Close closes the account returning the ammount that it contained.
func (account *Account) Close() (payout int64, ok bool) {
	account.mux.Lock()
	defer account.mux.Unlock()

	if account.closed {
		return 0, false
	}

	account.closed = true

	return account.balance, true
}

// Balance returns the current balance of the account.
func (account *Account) Balance() (balance int64, ok bool) {
	account.mux.Lock()
	defer account.mux.Unlock()

	if account.closed {
		return 0, false
	}

	return account.balance, true
}

// Deposit deposits the given amount. If the given value is negative it will be withdrawn.
func (account *Account) Deposit(ammount int64) (newBalance int64, ok bool) {
	account.mux.Lock()
	defer account.mux.Unlock()

	if account.closed || account.balance+ammount < 0 {
		return 0, false
	}

	account.balance += ammount

	return account.balance, true
}
