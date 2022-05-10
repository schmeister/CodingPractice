package account

import (
	"sync"
)

// Define the Account type here.
type Account struct {
	balance int64
	open    bool
	sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	account := Account{
		balance: amount,
		open:    true,
	}
	return &account
}

func (a *Account) Balance() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return 0, false
	}
	if (a.balance + amount) < 0 {
		return a.balance, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return 0, false
	}
	a.open = false
	payout := a.balance
	a.balance = 0
	return payout, true
}
