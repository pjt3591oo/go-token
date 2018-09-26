package receipt

import (
	"encoding/json"

	"../utils"
	"github.com/syndtr/goleveldb/leveldb"
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

const DB_DIRECTORY = "db/receipt"

func GetAllReceipt(account string) []Receipt {
	db, _ := leveldb.OpenFile(DB_DIRECTORY, nil)
	defer db.Close()

	var receipts []Receipt

	RootReceiptAsBytes, _ := db.Get([]byte(ROOT_RECEIPT+account+RECEIPT), nil)
	rootReceipt := RootReceipt{}
	json.Unmarshal(RootReceiptAsBytes, &rootReceipt)

	receiptAsBytes, _ := db.Get([]byte(NODE_RECEIPT+rootReceipt.ReceiptId+RECEIPT), nil)
	receipt := Receipt{}
	json.Unmarshal(receiptAsBytes, &receipt)

	for receipt.ReceiptId != "" {

		receipts = append(receipts, receipt)

		if receipt.NextReceiptId == "" {
			break
		}

		d, _ := db.Get([]byte(NODE_RECEIPT+receipt.NextReceiptId+RECEIPT), nil)
		json.Unmarshal(d, &receipt)
	}

	return receipts
}

func GetLastReceipt(account string) string {
	db, _ := leveldb.OpenFile(DB_DIRECTORY, nil)
	defer db.Close()

	receiptAsBytes, _ := db.Get([]byte(LAST_RECEIPT+account+RECEIPT), nil)

	return string(receiptAsBytes)
}
func GetRooteceipt(account string) string {
	db, _ := leveldb.OpenFile(DB_DIRECTORY, nil)
	defer db.Close()

	receiptAsBytes, _ := db.Get([]byte(ROOT_RECEIPT+account+RECEIPT), nil)

	return string(receiptAsBytes)
}

func GetSpecificReceipt(receiptId string) string {
	db, _ := leveldb.OpenFile(DB_DIRECTORY, nil)
	defer db.Close()

	receiptAsBytes, _ := db.Get([]byte(NODE_RECEIPT+receiptId+RECEIPT), nil)
	return string(receiptAsBytes)
}

func AddReceipt(txId string, toAccount string, status string, timestamp string) string {
	db, _ := leveldb.OpenFile(DB_DIRECTORY, nil)
	defer db.Close()

	receiptId := utils.Sha256(txId + timestamp + status)

	rr, _ := db.Get([]byte(ROOT_RECEIPT+toAccount+RECEIPT), nil) // 첫 번째 노드
	rootReceipt := RootReceipt{}
	json.Unmarshal([]byte(rr), &rootReceipt)

	lr, _ := db.Get([]byte(LAST_RECEIPT+toAccount+RECEIPT), nil) // 마지막 노드
	lastReceipt := LastReceipt{}
	json.Unmarshal([]byte(lr), &lastReceipt)

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
	db.Put([]byte(LAST_RECEIPT+toAccount+RECEIPT), lastReceiptAsBytes, nil)

	return receiptId
}
