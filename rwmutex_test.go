package belajar_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	rwMutex sync.RWMutex
	balance int
}

func (account *BankAccount) GetBalance() int {
	account.rwMutex.RLock()
	balance := account.balance
	account.rwMutex.RUnlock()
	return balance
}

func (account *BankAccount) AddBalance(amount int) {
	account.rwMutex.Lock()
    account.balance += amount
    account.rwMutex.Unlock()
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 1; i <= 1000; i++ {
		go func() {
			account.AddBalance(1)
			fmt.Println(account.GetBalance())
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(account.GetBalance())
}