package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
