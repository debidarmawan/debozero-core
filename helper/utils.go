package helper

import (
	"math/rand"
	"time"
)

func ContainsString(arr []string, item string) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}

	return false
}

func GenerateRandomString(length int) string {
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	b := make([]rune, length)

	for i := range b {
		b[i] = charset[randomizer.Intn(len(charset))]
	}

	return string(b)
}

func Contains[T string | int](arr []T, item T) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}

	return false
}
