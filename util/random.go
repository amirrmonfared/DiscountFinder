package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const priceString= "123456789"

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomPriceString(n int) string {
	var sb strings.Builder
	k := len(priceString)

	for i := 0; i < n; i++ {
		c := priceString[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomLink() string {
	return RandomString(70)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

