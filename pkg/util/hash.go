package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}
