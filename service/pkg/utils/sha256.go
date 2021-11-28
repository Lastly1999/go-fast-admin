package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// EncryptSh256 加密
func EncryptSh256(pass string) string {
	m := sha256.New()
	m.Write([]byte(pass))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
