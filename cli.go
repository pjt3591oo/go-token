package main

import (
	"flag"
	"fmt"
	"strconv"

	"./src/account"
	"./src/receipt"
)

type TransferFlag struct {
	from  string
	to    string
	value string
}

func main() {
	transfer := new(TransferFlag)

	accountKey := flag.String("createaccount", "", "create account")

	isTransfer := flag.String("istransfer", "0", "해당 옵션에 1을 전달할 경우 from, to, value 인자를 전달해야 함")
	flag.StringVar(&transfer.from, "from", "", "sender")
	flag.StringVar(&transfer.to, "to", "", "receiver")
	flag.StringVar(&transfer.value, "value", "", "token value")

	searchReceipt := flag.String("searchreceipt", "", "receipt 조회")

	flag.Parse()

	if *accountKey != "" {
		account, accountCreateErr := account.Allocation(*accountKey, 100, "user")

		if accountCreateErr != "" {
			fmt.Println("account 생성실패: ", accountCreateErr)
		} else {
			fmt.Println("used key: { ", *accountKey, " }")
			fmt.Println("created account: { ", account, " }")
		}
	} else if *isTransfer != "0" {
		from := transfer.from
		to := transfer.to
		value, _ := strconv.Atoi(transfer.value)

		a, b, c, d := account.Transfer(from, to, value)
		fmt.Println(a, b, c, d)
	} else if *searchReceipt != "" {
		// fmt.Println(*searchReceipt)
		r := receipt.GetReceipt(*searchReceipt)
		fmt.Println(r)
	}

}
