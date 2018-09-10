package main

import (
	"fmt"

	"./src/receipt"

	"./src/account"
)

func main() {
	var accounts [3]string

	accounts[0] = account.Allocation("card1", 100, "user")
	accounts[1] = account.Allocation("card2", 200, "user")
	accounts[2] = account.Allocation("card3", 300, "user")

	fmt.Println("BalanceOf : ", account.BalanceOf)
	// fmt.Println(token.Token())

	fmt.Println("========== token transfer test ==========")

	txId1, receiptIdIn1, receiptIdOut1, err1 := account.BalanceOf[accounts[0]].Transfer(accounts[1], 30)

	if err1 {
		fmt.Println("test case 1. Filed")
	} else {
		fmt.Println("txId: ", txId1)
		// fmt.Println("test case 1. transfer retulr : ", txId1, account.BalanceOf[accounts[0]], account.BalanceOf[accounts[1]])
	}

	txId2, receiptIdIn2, receiptIdOut2, err2 := account.BalanceOf[accounts[0]].Transfer(accounts[2], 10)

	if err2 {
		fmt.Println("test case 2. Filed: amount lack")
	} else {
		fmt.Println("txId: ", txId2)
		// fmt.Println("test case 2. transfer retulr : ", txId2, account.BalanceOf[accounts[0]], account.BalanceOf[accounts[1]])
	}

	txId3, receiptIdIn3, receiptIdOut3, err3 := account.BalanceOf[accounts[0]].Transfer(accounts[1], 20)

	if err3 {
		fmt.Println("test case 3. Filed: amount lack")
	} else {
		fmt.Println("txId: ", txId3)
		// fmt.Println("test case 3. transfer retulr : ", txId3, account.BalanceOf[accounts[0]], account.BalanceOf[accounts[1]])
	}

	fmt.Println(" ----------- Inresult ----------- ")
	fmt.Println(" \n", receiptIdIn1, receipt.GetReceipt(receiptIdIn1))
	fmt.Println(" \n", receiptIdIn2, receipt.GetReceipt(receiptIdIn2))
	fmt.Println(" \n", receiptIdIn3, receipt.GetReceipt(receiptIdIn3))

	fmt.Println(" ----------- Outresult ----------- ")
	fmt.Println(" \n", receiptIdOut1, receipt.GetReceipt(receiptIdOut1))
	fmt.Println(" \n", receiptIdOut2, receipt.GetReceipt(receiptIdOut2))
	fmt.Println(" \n", receiptIdOut3, receipt.GetReceipt(receiptIdOut3))

	fmt.Println("----------- result -----------")
	fmt.Println(receipt.ShowAccountReceipt(accounts[2]))

	// isSuccess2 := token.Transfer(accounts[0], accounts[1], 200)
	// fmt.Println("1st transfer retulr : ", isSuccess2, account.BalanceOf)
}
