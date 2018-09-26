package transaction

import (
	"encoding/json"

	"../utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type Transaction struct {
	TxId      string
	From      string
	To        string
	Value     int
	Timestamp string
}

const DB_DIRECTORY = "db/transaction"

func GetTransaction(txId string) (string, string) {
	if utils.Sha256ValidCheck(txId) {
		return "", "transaction ID가 옳바르지 않습니다."
	}

	db, _ := leveldb.OpenFile(DB_DIRECTORY, nil)
	defer db.Close()

	transactionAsBytes, _ := db.Get([]byte(txId), nil)

	return string(transactionAsBytes), ""
}

func CreateTransaction(from string, to string, value int, timestamp string) (string, bool) {
	db, err := leveldb.OpenFile(DB_DIRECTORY, nil)
	defer db.Close()

	if err != nil {
		return "", true
	}

	txId := utils.Sha256(from + to + string(value) + timestamp)

	tx := Transaction{
		TxId:      txId,
		From:      from,
		To:        to,
		Value:     value,
		Timestamp: timestamp,
	}

	transactionAsBytes, _ := json.Marshal(tx)
	_ = db.Put([]byte(txId), transactionAsBytes, nil)

	return txId, false
}
