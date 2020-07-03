package util

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	_ , err:= m.Write([]byte(value))
	if err != nil{
		log.Fatal("md5 write error")
	}
	return hex.EncodeToString(m.Sum(nil))
}
