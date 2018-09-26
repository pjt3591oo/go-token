package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"./src/account"
	"./src/receipt"
	"./src/transaction"

	log "github.com/withmandala/go-log" // https://github.com/Mandala/go-log
)

type TransferFlag struct {
	from  string
	to    string
	value string
}

func main() {
	logger := log.New(os.Stderr).WithColor().WithDebug().WithoutTimestamp()
	transfer := new(TransferFlag)

	// option 처리
	accountKey := flag.String("createaccount", "", "create account")

	isTransfer := flag.String("istransfer", "0", "해당 옵션에 1을 전달할 경우 from, to, value 인자를 전달해야 함")
	flag.StringVar(&transfer.from, "from", "", "sender")
	flag.StringVar(&transfer.to, "to", "", "receiver")
	flag.StringVar(&transfer.value, "value", "", "token value")

	searchReceipt := flag.String("searchreceipt", "", "receipt 조회")

	accountReceipt := flag.String("accountreceipt", "", "account의 receipt내역, root, last 조회")

	searchTx := flag.String("searchtx", "", "transaction 조회")
	searchAccount := flag.String("searchaccount", "", "account 조회")

	// option 파싱
	flag.Parse()

	if *accountKey != "" {
		address, accountCreateErr := account.Allocation(*accountKey, 100, "user")

		if accountCreateErr != "" {
			logger.Error("account 생성실패: ", accountCreateErr)
		} else {
			logger.Info("used key: { ", *accountKey, " }")
			logger.Info("created account address: { ", address, " }")
		}
	} else if *isTransfer != "0" {
		from := transfer.from
		to := transfer.to
		value, _ := strconv.Atoi(transfer.value)

		txId, _, _, txErr := account.Transfer(from, to, value)

		if txErr != "" {
			logger.Error(txErr)
		} else {
			logger.Info("transactionId: ", txId)
		}

	} else if *searchReceipt != "" {
		// fmt.Println(*searchReceipt)
		r := receipt.GetSpecificReceipt(*searchReceipt)
		logger.Info(r)
	} else if *accountReceipt != "" {
		rr := receipt.GetRooteceipt(*accountReceipt)
		lr := receipt.GetLastReceipt(*accountReceipt)
		ar := receipt.GetAllReceipt(*accountReceipt)

		logger.Info("ROOT RECEIPT : ", rr)
		logger.Info("LAST RECEIPT : ", lr)
		logger.Info("All RECEIPT ", "(", len(ar), "): ", ar)
	} else if *searchTx != "" {
		tx, txErr := transaction.GetTransaction(*searchTx)

		if txErr != "" {
			logger.Error(txErr)
		} else {
			logger.Info(tx)
		}
	} else if *searchAccount != "" {
		ac, accountErr := account.GetAccountInfo(*searchAccount)

		if accountErr != "" {
			logger.Error(accountErr)
		} else {
			logger.Info(ac)
		}
	} else {
		msg := ` 
	-accountreceipt string
			account의 receipt내역, root, last 조회
	-createaccount string
			create account
	-from string
			sender
	-istransfer string
			해당 옵션에 1을 전달할 경우 from, to, value 인자를 전달해야 함
	-searchaccount string
			account 조회
	-searchreceipt string
			receipt 조회
	-searchtx string
			transaction 조회
	-to string
			receiver
	-value string
			token value
		`

		fmt.Println(msg)

	}

}
