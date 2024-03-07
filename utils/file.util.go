package utils

import (
	"math/rand"
	"os"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)

	if err != nil {
		return err
	}

	return nil
}
