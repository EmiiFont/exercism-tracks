package account

import "sync"

// Define the Account type here.
type Account struct {
	mu      sync.Mutex
	balance int64
	active  bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, active: true}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.active {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.active || (amount < 0 && a.balance+amount < 0) {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.active {
		return 0, false
	}
	current := a.balance
	a.balance = 0
	a.active = false
	return current, true
}
