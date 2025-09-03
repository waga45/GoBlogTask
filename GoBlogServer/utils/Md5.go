package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const key = "hahaGo90"

/*
*
md5
*/
func Md5String(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum([]byte(key)))
}
