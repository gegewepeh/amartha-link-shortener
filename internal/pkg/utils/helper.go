package utils

import (
	"math/rand"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandString generate random string
func RandString(n int, forceUpper bool) string {
	rand.Seed(time.Now().UnixNano())

	var str string
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
		str = string(b)
	}

	if forceUpper {
		str = strings.ToUpper(str)
	}

	return str
}
