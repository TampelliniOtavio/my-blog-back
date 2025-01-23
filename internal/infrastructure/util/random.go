package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var currRand *rand.Rand

func init() {
	currRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[currRand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail(nameSize int, domainSize int) string {
	return RandomString(nameSize) + "@" + RandomString(domainSize) + ".com"
}
