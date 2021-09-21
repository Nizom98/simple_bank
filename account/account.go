package account

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
)

var StoreData Store = Store{MutexAccounts: &sync.Mutex{}, MutexNextID: &sync.Mutex{}} //

func NewAccount(balance int64) (newAccount Account) {
	StoreData.MutexNextID.Lock()
	newAccount = Account{ID: StoreData.NextID, Balance: balance, MutexBalance: &sync.Mutex{}}
	StoreData.NextID++
	StoreData.MutexNextID.Unlock()

	StoreData.MutexAccounts.Lock()
	defer StoreData.MutexAccounts.Unlock()
	StoreData.Accounts = append(StoreData.Accounts, newAccount)
	return
}

func GetAccountBalance(ID int64) (balance int64, err error) {
	for _, account := range StoreData.Accounts {
		if account.ID == ID {
			return account.Balance, nil
		}
	}
	return 0, errors.New("аккаунт с таким идетификатором не найден")
}

func TransferBalance(fromID, toID, sum int64) (err error) {
	if fromID == toID {
		return errors.New("передача средств в рамках одного аккаунта не допустима")
	} else if sum < 1 {
		return errors.New("минимальная сумма списания равна 1 руб")
	}
	StoreData.MutexAccounts.Lock()
	defer StoreData.MutexAccounts.Unlock()
	var fromIDX, toIDX int = -1, -1
	for i, account := range StoreData.Accounts {
		if account.ID == fromID {
			fromIDX = i
		} else if account.ID == toID {
			toIDX = i
		} else if fromIDX != -1 && toIDX != -1 {
			break
		}
	}
	if fromIDX == -1 || toIDX == -1 {
		return errors.New("один из аккаунтов не найден")
	} else if StoreData.Accounts[fromIDX].Balance < sum {
		return errors.New("баланс аккаунта-источника меньше, чем сумма списания")
	}
	StoreData.Accounts[fromIDX].Balance -= sum
	StoreData.Accounts[toIDX].Balance += sum
	return nil

}

func rb2map(r *http.Request) (requestBody map[string]interface{}) {
	var body map[string]interface{}
	json.NewDecoder(r.Body).Decode(&body)
	return body
}
