package utils

import (
	"math/rand"
	"time"
)

const lower_char = "abcdefghijklmnopqrstuvwxyz"
const higher_char = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const number = "0123456789"
const charset = lower_char + higher_char + number

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
