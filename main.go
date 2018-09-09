package main

import (
	"fmt"

	"./src/account"
)

func main() {
	account.Allocation(100)
	account.Allocation(200)
	account.Allocation(300)

	fmt.Println(account.BalanceOf)
}
