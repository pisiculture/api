package util

import (
	"crypto/sha256"
)

func Crypto(text string) (string, error) {
	res := sha256.Sum256([]byte(text))
	return string(res[:]), nil
}
