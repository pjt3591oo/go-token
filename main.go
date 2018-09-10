package main

import (
	"fmt"

	"./src/account"
	"./src/token"
)

func main() {
	var accounts [3]string

	accounts[0] = account.Allocation(100)
	accounts[1] = account.Allocation(200)
	accounts[2] = account.Allocation(300)

	fmt.Println("BalanceOf : ", account.BalanceOf)

	fmt.Println(token.Token())

	fmt.Println("========== token transfer test ==========")

	isSuccess1 := token.Transfer(accounts[0], accounts[1], 50)
	fmt.Println("1st transfer retulr : ", isSuccess1, account.BalanceOf)

	isSuccess2 := token.Transfer(accounts[0], accounts[1], 200)
	fmt.Println("1st transfer retulr : ", isSuccess2, account.BalanceOf)
}
