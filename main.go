package main

import (
	"fmt"
	// "go/token"

	"./src/account"
	"./src/receipt"
)

func main() {
	var accounts [3]string

	accounts[0] = account.Allocation("card1", 100, "user")
	accounts[1] = account.Allocation("card2", 200, "user")
	accounts[2] = account.Allocation("card3", 300, "user")

	fmt.Println("BalanceOf : ", account.BalanceOf)
	// fmt.Println(token.Token())

	fmt.Println("========== token transfer test ==========")

	txId1, receiptId1, err1 := account.BalanceOf[accounts[0]].Transfer(accounts[1], 30)

	if err1 {
		fmt.Println("test case 1. Filed")
	} else {
		fmt.Println("txId: ", txId1)
		// fmt.Println("test case 1. transfer retulr : ", txId1, account.BalanceOf[accounts[0]], account.BalanceOf[accounts[1]])
	}

	txId2, receiptId2, err2 := account.BalanceOf[accounts[0]].Transfer(accounts[1], 10)

	if err2 {
		fmt.Println("test case 2. Filed: amount lack")
	} else {
		fmt.Println("txId: ", txId2)
		// fmt.Println("test case 2. transfer retulr : ", txId2, account.BalanceOf[accounts[0]], account.BalanceOf[accounts[1]])
	}

	txId3, receiptId3, err3 := account.BalanceOf[accounts[0]].Transfer(accounts[1], 20)

	if err3 {
		fmt.Println("test case 3. Filed: amount lack")
	} else {
		fmt.Println("txId: ", txId3)
		// fmt.Println("test case 3. transfer retulr : ", txId3, account.BalanceOf[accounts[0]], account.BalanceOf[accounts[1]])
	}

	fmt.Println(" ----------- result ----------- ")
	fmt.Println(" \n", receiptId1, receipt.GetReceipt(receiptId1))
	fmt.Println(" \n", receiptId2, receipt.GetReceipt(receiptId2))
	fmt.Println(" \n", receiptId3, receipt.GetReceipt(receiptId3))
	// isSuccess2 := token.Transfer(accounts[0], accounts[1], 200)
	// fmt.Println("1st transfer retulr : ", isSuccess2, account.BalanceOf)
}
