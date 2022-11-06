package shop

import (
	"math/rand"
	"strings"
	"time"
)

const lengthID = 10
const charset = "abcdefghijklmnopqrstuvwxyz1234567890"

func GenerateID() string {
	rand.Seed(time.Now().UnixNano())

	var tokens []string
	for i := 0; i <= lengthID; i++ {
		index := rand.Intn(len(charset))
		tokens = append(tokens, string(charset[index]))
	}

	return strings.Join(tokens, "")
}
