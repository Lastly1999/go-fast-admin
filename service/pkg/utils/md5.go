package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// DeEncodeMD5 md5
func DeEncodeMD5(value string) string {
	dst := make([]byte, hex.DecodedLen(len([]byte(value))))
	n, err := hex.Decode(dst, []byte(value))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	return ""
}
