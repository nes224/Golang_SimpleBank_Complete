package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklnmopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano()) // as rand.Seed() expect an int64 as input, .UnixNano() convert the time to unix nano
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1) // 0->max-min
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i< n; i ++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c) // We call sb.WriteByte() to write that character c to the string builder
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney gneerates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

