package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random generate a integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random generate a string of lenth n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

// Random balance on account
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Random currency code
func RandomCurrency() string {
	currecies := []string{EUR, USD, UAH}
	n := len(currecies)

	return currecies[rand.Intn(n)]
}
