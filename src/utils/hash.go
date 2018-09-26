package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(message string) string {
	hash := sha256.New()

	hash.Write([]byte(message))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}
