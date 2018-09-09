package account

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

var BalanceOf = make(map[string]int)

func Sha256(message string) string {
	hash := sha256.New()

	hash.Write([]byte(message))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}

func NewAccount() string {
	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	converted := Sha256(string(timestamp))

	return converted
}

func Allocation(balance int) {
	BalanceOf[NewAccount()] = balance
}
