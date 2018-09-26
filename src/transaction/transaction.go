package transaction

import (
	"encoding/json"

	"../utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type transaction struct {
	TxId      string
	From      string
	To        string
	Value     int
	Timestamp string
}

var Transaction = make(map[string]transaction)

func CreateTransaction(from string, to string, value int, timestamp string) (string, bool) {
	db, err := leveldb.OpenFile("db/transaction", nil)
	defer db.Close()

	if err != nil {
		return "", true
	}

	txId := utils.Sha256(from + to + string(value) + timestamp)

	tx := transaction{
		TxId:      txId,
		From:      from,
		To:        to,
		Value:     value,
		Timestamp: timestamp,
	}

	Transaction[txId] = tx

	transactionAsBytes, _ := json.Marshal(tx)
	_ = db.Put([]byte(txId), transactionAsBytes, nil)

	return txId, false
}
