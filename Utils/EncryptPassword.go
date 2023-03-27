package Utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 加密盐
const secret = "yiren"

// 密码加密
func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
