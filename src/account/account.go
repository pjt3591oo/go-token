package account

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	"../receipt"
	"../transaction"
)

type Account struct {
	Value   int    // 토큰 량
	Kind    string // user, store 구분
	Address string
}

var BalanceOf = make(map[string]*Account)

func Sha256(message string) string {
	hash := sha256.New()

	hash.Write([]byte(message))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}

func NewAccount(cardId string) string {
	// timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	converted := Sha256(cardId)

	return converted
}

func Allocation(cardId string, balance int, kind string) string {
	createdAccount := NewAccount(cardId)
	BalanceOf[createdAccount] = &Account{balance, kind, createdAccount}

	return createdAccount
}

func (account *Account) Transfer(to string, amount int) (string, string, string, bool) {
	if account.Value < amount {
		return "", "", "", true
	}

	account.Value -= amount
	BalanceOf[to].Value += amount

	timestamp := strconv.Itoa(int(time.Now().UnixNano()))

	txId := transaction.CreateTransaction(account.Address, to, amount, timestamp)

	// receiptId := receipt.AddReceipt(txId, timestamp)
	receiptIdOut := receipt.AddReceipt(txId, account.Address, "OUT", timestamp)
	receiptIdIn := receipt.AddReceipt(txId, to, "IN", timestamp)

	return txId, receiptIdIn, receiptIdOut, false
}
