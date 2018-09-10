package token

import (
	"../account"
)

var Name = "test token"
var Symbol = "tt"
var TotalAmount = 100000000

func Token() (string, string, int) {
	return Name, Symbol, TotalAmount
}

func Transfer(to string, from string, amount int) bool {

	if account.BalanceOf[from] < amount {
		return false
	}

	account.BalanceOf[from] -= amount
	account.BalanceOf[to] += amount

	return true
}

func AddPublish(amount int) {
	TotalAmount += amount
}
