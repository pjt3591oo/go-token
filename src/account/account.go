package account

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"../receipt"
	"../transaction"
	"../utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type Account struct {
	Value   int    // 토큰 량
	Kind    string // user, store 구분
	Address string
}

var BalanceOf = make(map[string]*Account)

func isAddress(address string) bool {
	db, _ := leveldb.OpenFile("db/account", nil)
	defer db.Close()

	data, _ := db.Get([]byte(address), nil)
	account := Account{}

	json.Unmarshal(data, &account)

	if account.Address != "" {
		return true
	}

	return false
}

func newAccount(cardId string) (string, bool) {
	// timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	converted := utils.Sha256(cardId)
	if isAddress(converted) {
		return "", true
	}

	return converted, false
}

func Allocation(cardId string, balance int, kind string) (string, string) {
	createdAccount, isAccount := newAccount(cardId)

	if isAccount {
		return "", "이미 account가 존재합니다."
	}

	BalanceOf[createdAccount] = &Account{balance, kind, createdAccount}

	db, err := leveldb.OpenFile("db/account", nil)
	defer db.Close()

	if err != nil {
		return "", "디비 연결중 문제발생"
	}

	createdAccountAsBytes, _ := json.Marshal(BalanceOf[createdAccount])
	_ = db.Put([]byte(createdAccount), createdAccountAsBytes, nil)

	return createdAccount, ""
}

func AccountInvalidCheck(address string) bool {
	if len(address) != 64 {
		return true
	}

	return false
}

func accountCheck(from string, to string) (bool, string) {
	// 1. address 길이확이
	// 2. address 존재확인
	if AccountInvalidCheck(from) == true {
		fmt.Println("invalie account")
		return true, "invalid from account"
	}
	if AccountInvalidCheck(to) == true {
		fmt.Println("invalie account")
		return true, "invalid to account"
	}

	if isAddress(from) == false {
		return true, "from address가 존재하지 않습니다."
	}

	if isAddress(to) == false {
		return true, "to address가 존재하지 않습니다."
	}

	return false, ""
}

func Transfer(from string, to string, amount int) (string, string, string, string) {
	invalidChecked, msg := accountCheck(from, to)
	if invalidChecked {
		return "", "", "", msg
	}

	db, _ := leveldb.OpenFile("db/account", nil)
	defer db.Close()

	// 보내는 account 데이터 가져오기
	fromData, _ := db.Get([]byte(from), nil)
	fromAccount := Account{}

	json.Unmarshal(fromData, &fromAccount)

	if fromAccount.Value < amount {
		return "", "", "", "잔액이 부족합니다."
	}

	// 보내는 account value 감소
	fromAccount.Value -= amount

	// 보내는 account 데이터 업데이트
	fromDataAsBytes, _ := json.Marshal(fromAccount)
	db.Put([]byte(from), []byte(fromDataAsBytes), nil)

	// 받는 account 데이터 가져오기
	toData, _ := db.Get([]byte(to), nil)
	toAccount := Account{}

	json.Unmarshal(toData, &toAccount)

	// 받는 account value 감소
	toAccount.Value += amount

	// 받는 account 데이터 업데이트
	toDataAsBytes, _ := json.Marshal(toAccount)
	db.Put([]byte(to), []byte(toDataAsBytes), nil)

	timestamp := strconv.Itoa(int(time.Now().UnixNano()))

	txId, _ := transaction.CreateTransaction(from, to, amount, timestamp)

	receiptIdOut := receipt.AddReceipt(txId, from, "OUT", timestamp)
	receiptIdIn := receipt.AddReceipt(txId, to, "IN", timestamp)

	return txId, receiptIdIn, receiptIdOut, ""
}
