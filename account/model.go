package account

import (
	"sync"
)

type Store struct {
	MutexAccounts *sync.Mutex
	Accounts      []Account
	MutexNextID   *sync.Mutex
	NextID        int64
}

type Account struct {
	ID           int64       `json:"id"`
	MutexBalance *sync.Mutex `json:"-"`
	Balance      int64       `json:"balance"`
}
