package main

import (
	"os"

	"github.com/withmandala/go-log" // https://github.com/Mandala/go-log
)

func main() {

	logger := log.New(os.Stderr).WithColor().WithDebug().WithoutTimestamp()

	// var accounts [3]string

	logger.Trace("========== account create test ==========")
	// accounts[0], _ = account.Allocation("card1", 100, "user")
	// accounts[1], _ = account.Allocation("card2", 200, "user")
	// accounts[2], _ = account.Allocation("card3", 300, "user")

	// fmt.Println("BalanceOf : ", account.BalanceOf)

	logger.Trace("========== token transfer success test ==========")

	// txId1, receiptIdIn1, receiptIdOut1, err1 := account.BalanceOf[accounts[0]].Transfer(accounts[1], 30)

	// if err1 {
	// 	logger.Warn("test case 1. Filed")
	// } else {
	// 	logger.Info("txId: ", txId1, ", err1: ", err1)
	// }

	// txId2, receiptIdIn2, receiptIdOut2, err2 := account.BalanceOf[accounts[2]].Transfer(accounts[0], 10)

	// if err2 {
	// 	logger.Warn("test case 2. Filed: amount lack")
	// } else {
	// 	logger.Info("txId: ", txId2, ", err2: ", err2)
	// }

	// txId3, receiptIdIn3, receiptIdOut3, err3 := account.BalanceOf[accounts[0]].Transfer(accounts[2], 20)

	// if err3 {
	// 	logger.Warn("test case 3. Filed: amount lack")
	// } else {
	// 	logger.Info("txId: ", txId3, ", err3: ", err3)
	// }

	logger.Trace("========== token transfer failed test ==========")

	// txId4, _, _, err4 := account.BalanceOf[accounts[0]].Transfer("5dd297459441", 30)
	// fmt.Println("txId: ", txId4, ", err: ", err4)

	logger.Trace(" ----------- Inresult ----------- ")
	// fmt.Println(" \n", receipt.GetReceipt(receiptIdIn1))
	// fmt.Println(" \n", receipt.GetReceipt(receiptIdIn2))
	// fmt.Println(" \n", receipt.GetReceipt(receiptIdIn3))

	logger.Trace(" ----------- Outresult ----------- ")
	// fmt.Println(" \n", receipt.GetReceipt(receiptIdOut1))
	// fmt.Println(" \n", receipt.GetReceipt(receiptIdOut2))
	// fmt.Println(" \n", receipt.GetReceipt(receiptIdOut3))

	logger.Trace("----------- account0 receipt result -----------")
	// receipt.ShowAccountReceipt(accounts[0])

	logger.Trace("----------- account1 receipt result -----------")
	// receipt.ShowAccountReceipt(accounts[1])

	logger.Trace("----------- account2 receipt result -----------")
	// receipt.ShowAccountReceipt(accounts[2])

}
