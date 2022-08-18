package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Encoder 密码加密
func Encoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
