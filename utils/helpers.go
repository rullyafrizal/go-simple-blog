package utils

import (
	"math/rand"
	"net/mail"
	"os"
	"time"
)

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func IsValidEmail(v string) bool {
	_, err := mail.ParseAddress(v)

	return err == nil
}

func StringifyErrors(err ...error) []string {
	var errors []string

	for _, e := range err {
		errors = append(errors, e.Error())
	}

	return errors
}

func GenerateRandomString(length int) string {
	return GenerateRandomStringWithCharset(length, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

func GenerateRandomStringWithCharset(length int, charset string) string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
