package main

import (
	"Proj1/account"
	"net/http"
)

func main() {
	http.HandleFunc("/account/new", account.HandlerNewAccount) //
	http.HandleFunc("/account/balance", account.HandlerGetBalance)
	http.HandleFunc("/account/transfer_balance", account.HandlerTransferBalance)
	err := http.ListenAndServe(":4000", nil)
	panic(err)
}
