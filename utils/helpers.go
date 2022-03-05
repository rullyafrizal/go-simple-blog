package utils

import (
	"net/mail"
	"os"
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