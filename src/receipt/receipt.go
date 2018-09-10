package receipt

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type receipt struct {
	ReceiptId     string `json:"name"`
	TxId          string `json:"txId"`
	NextReceiptId string `json:"nextReceiptId"`
	PrevRecieptId string `json:"prevRecieptId"`
	LastRecieptId string `json:"lastRecieptId"`
	Status        string `json:"status"`
}

var AccountLastReceiptId = make(map[string]string)

var Receipt = make(map[string]*receipt)       // 각각의 결과를 하나하나 키로 저장
var AccountReceipt = make(map[string]receipt) // 각각의 결과를 유저에게 리스트 형태로 저장

func Sha256(message string) string {
	hash := sha256.New()

	hash.Write([]byte(message))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}

func GetReceipt(receiptId string) string {
	return GetJson(*Receipt[receiptId])
}

func ShowAccountReceipt(toAccount string) receipt {
	// startReceiptId = AccountReceipt[toAccount].ReceiptId
	i := 0
	for receiptId := AccountReceipt[toAccount].ReceiptId; i < 2; i++ {
		fmt.Println(Receipt[receiptId])
		receiptId = Receipt[receiptId].NextReceiptId

	}

	return AccountReceipt[toAccount]
}

func GetJson(data receipt) string {
	U, _ := json.Marshal(data)
	var usr receipt
	json.Unmarshal(U, &usr)

	return string(U)
}

func (PrevReceipt *receipt) SetNextReceiptId(nextReceiptId string) {
	PrevReceipt.NextReceiptId = nextReceiptId
}

func AddReceipt(txId string, toAccount string, status string, timestamp string) string {
	receiptId := Sha256(txId + timestamp + status)
	// fromAccount := transaction.Transaction[txId].From

	// 각각의 마지막 내역에 추가된 내역의 ID를 nextReceiptId에 넣는다
	toPrevRecieptId := AccountLastReceiptId[toAccount]

	fmt.Println("\n toPrevRecieptId : ", toPrevRecieptId)
	fmt.Println(GetJson(AccountReceipt[toAccount]))

	Receipt[receiptId] = &receipt{
		ReceiptId:     receiptId,
		TxId:          txId,
		NextReceiptId: "",
		PrevRecieptId: toPrevRecieptId,
		LastRecieptId: receiptId,
		Status:        status,
	}

	d := GetJson(AccountReceipt[toAccount])

	if d == "{\"name\":\"\",\"txId\":\"\",\"nextReceiptId\":\"\",\"prevRecieptId\":\"\",\"lastRecieptId\":\"\",\"status\":\"\"}" {
		AccountReceipt[toAccount] = *Receipt[receiptId]
	} else {
		Receipt[toPrevRecieptId].SetNextReceiptId(Receipt[receiptId].LastRecieptId)
	}
	AccountLastReceiptId[toAccount] = receiptId

	return receiptId
}
