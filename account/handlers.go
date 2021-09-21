package account

import (
	"encoding/json"
	"net/http"
)

func HandlerNewAccount(w http.ResponseWriter, r *http.Request) { //Обработчик запроса /
	body := rb2map(r)
	balanceF64, ok := body["balance"].(float64)
	if !ok {
		printAnswer(w, http.StatusBadRequest, nil, nil, "Недопустимые параметры")
		return
	}
	newAccount := NewAccount(int64(balanceF64))
	printAnswer(w, http.StatusOK, newAccount, nil, "")

}

func HandlerGetBalance(w http.ResponseWriter, r *http.Request) { //Обработчик запроса /
	body := rb2map(r)
	accountIDF64, ok := body["id"].(float64)
	if !ok {
		printAnswer(w, http.StatusBadRequest, nil, nil, "Недопустимые параметры")
		return
	}
	balance, err := GetAccountBalance(int64(accountIDF64))
	printAnswer(w, http.StatusOK, balance, err, "")

}

func HandlerTransferBalance(w http.ResponseWriter, r *http.Request) { //Обработчик запроса /
	body := rb2map(r)
	fromID, fromOk := body["from_id"].(float64)
	toID, toOk := body["to_id"].(float64)
	sum, sumOk := body["sum"].(float64)
	if !fromOk || !toOk || !sumOk {
		printAnswer(w, http.StatusBadRequest, nil, nil, "Недопустимые параметры")
		return
	}
	if err := TransferBalance(int64(fromID), int64(toID), int64(sum)); err != nil {
		printAnswer(w, http.StatusInternalServerError, nil, err, "")
		return
	}
	printAnswer(w, http.StatusOK, "Успешно", nil, "")
}

func printAnswer(w http.ResponseWriter, httpCode int, data interface{}, err error, altErr string) error {
	res := make(map[string]interface{})
	res["data"] = data
	res["ok"] = true
	isAltErr := len(altErr) > 0
	isErr := err != nil
	if isAltErr || isErr {
		res["ok"] = false
		if isAltErr {
			res["message"] = altErr
		} else if isErr {
			res["message"] = err.Error()
		}
	}
	w.WriteHeader(httpCode)
	return json.NewEncoder(w).Encode(res)
}
