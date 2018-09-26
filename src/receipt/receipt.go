package receipt

import (
	"encoding/json"
	"fmt"
	"os"

	"../utils"
	"github.com/syndtr/goleveldb/leveldb"
	log "github.com/withmandala/go-log"
)

type Receipt struct {
	ReceiptId     string `json:"receiptId"`
	TxId          string `json:"txId"`
	NextReceiptId string `json:"nextReceiptId"`
	PrevReceiptId string `json:"prevReceiptId"`
	Status        string `json:"status"`
}

type RootReceipt struct {
	ReceiptId string `json:"receiptId"`
}

type LastReceipt struct {
	ReceiptId string `json:"receiptId"`
}

const ROOT_RECEIPT = "ROOT_"
const NODE_RECEIPT = "NODE_"
const LAST_RECEIPT = "LAST_"
const RECEIPT = "RECEIPT"

var AccountLastReceiptId = make(map[string]string)

var logger = log.New(os.Stderr).WithColor().WithoutTimestamp()

// func ShowAccountReceipt(toAccount string) receipt {
// 	// startReceiptId = AccountReceipt[toAccount].ReceiptId
// 	i := 0
// 	for receiptId := AccountReceipt[toAccount].ReceiptId; true; i++ {
// 		logger.Info(GetReceipt(receiptId))
// 		receiptId = Receipt[receiptId].NextReceiptId

// 		if receiptId == "" {
// 			break
// 		}

// 	}

// 	return AccountReceipt[toAccount]
// }

func GetReceipt(receiptId string) string {
	db, _ := leveldb.OpenFile("db/receipt", nil)
	defer db.Close()

	// receipt := Receipt{}

	receiptAsBytes, _ := db.Get([]byte(NODE_RECEIPT+receiptId+RECEIPT), nil)
	// json.Unmarshal(receiptAsBytes, &receipt)

	return string(receiptAsBytes)
}

func AddReceipt(txId string, toAccount string, status string, timestamp string) string {
	db, _ := leveldb.OpenFile("db/receipt", nil)
	defer db.Close()

	receiptId := utils.Sha256(txId + timestamp + status)

	rr, _ := db.Get([]byte(ROOT_RECEIPT+toAccount+RECEIPT), nil) // 첫 번째 노드
	rootReceipt := RootReceipt{}
	json.Unmarshal([]byte(rr), &rootReceipt)

	lr, _ := db.Get([]byte(LAST_RECEIPT+toAccount+RECEIPT), nil) // 마지막 노드
	lastReceipt := LastReceipt{}
	json.Unmarshal([]byte(lr), &lastReceipt)

	e, _ := json.Marshal(lastReceipt)
	fmt.Println("*lstReceipt*: ", string(e))

	receipt := Receipt{
		ReceiptId:     receiptId,
		TxId:          txId,
		NextReceiptId: "",
		PrevReceiptId: lastReceipt.ReceiptId,
		Status:        status,
	}

	receiptAsBytes, _ := json.Marshal(receipt)

	if (rootReceipt == (RootReceipt{})) == true {
		rootReceipt.ReceiptId = receiptId
		rootReceiptAsBytes, _ := json.Marshal(receipt)

		fmt.Println("**root receipt id**", string(rootReceiptAsBytes))
		db.Put([]byte(ROOT_RECEIPT+toAccount+RECEIPT), rootReceiptAsBytes, nil)
		db.Put([]byte(NODE_RECEIPT+receiptId+RECEIPT), receiptAsBytes, nil)
	} else {
		prevReceiptId := lastReceipt.ReceiptId
		prevReceipt, _ := db.Get([]byte(NODE_RECEIPT+prevReceiptId+RECEIPT), nil)
		pr := Receipt{}
		json.Unmarshal([]byte(prevReceipt), &pr)

		pr.NextReceiptId = receiptId
		prAsBytes, _ := json.Marshal(pr)
		db.Put([]byte(NODE_RECEIPT+prevReceiptId+RECEIPT), prAsBytes, nil)

		db.Put([]byte(NODE_RECEIPT+receiptId+RECEIPT), receiptAsBytes, nil)
	}

	lastReceipt.ReceiptId = receiptId
	lastReceiptAsBytes, _ := json.Marshal(lastReceipt)
	fmt.Println("***last receipt id***", string(lastReceiptAsBytes))
	db.Put([]byte(LAST_RECEIPT+toAccount+RECEIPT), lastReceiptAsBytes, nil)

	return receiptId
}
