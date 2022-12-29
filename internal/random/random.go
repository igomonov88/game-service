package random

import (
	"math/rand"
	"time"
)

const (
	charset         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "abcdefghijklmnopqrstuvwxyz"
	defaultMaxValue = 5
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func String(length int) string {
	return stringWithCharset(length, charset)
}

func Int(maxInt int) int {
	if maxInt <= 0 {
		maxInt = defaultMaxValue
	}
	result := random.Intn(maxInt)
	if result == 0 {
		return 1
	}
	return result
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)
}
