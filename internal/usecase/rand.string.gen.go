package usecase

import (
	"math/rand"
)

type generateRandomString func(lenth int) string

func newRandomStringGenerator(randSeed int64) generateRandomString {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rs := rand.New(rand.NewSource(randSeed))
	return func(lenth int) string {
		b := make([]byte, lenth)
		for i := range b {
			b[i] = charset[rs.Intn(len(charset))]
		}
		return string(b)
	}
}
