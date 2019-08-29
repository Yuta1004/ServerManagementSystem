package common

import (
	"time"
	"math/rand"
)

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenPassphrase : 指定文字数のランダムな文字列を生成する
func GenPassphrase(n int) string {
	rand.Seed(time.Now().UnixNano())
	passphrase := make([]byte, n)
	for idx := 0; idx < n; idx ++ {
		alpIdx := rand.Intn(len(alphabets))
		passphrase[idx] = alphabets[alpIdx]
	}
	return string(passphrase)
}