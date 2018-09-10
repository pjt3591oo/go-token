package transaction

import (
	"crypto/sha256"
	"encoding/hex"
)

type transaction struct {
	TxId      string
	From      string
	To        string
	Value     int
	Timestamp string
}

var Transaction = make(map[string]transaction)

func Sha256(message string) string {
	hash := sha256.New()

	hash.Write([]byte(message))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}

func CreateTransaction(from string, to string, value int, timestamp string) string {

	txId := Sha256(from + to + string(value) + timestamp)

	tx := transaction{
		TxId:      txId,
		From:      from,
		To:        to,
		Value:     value,
		Timestamp: timestamp,
	}

	Transaction[txId] = tx
	return txId
}
