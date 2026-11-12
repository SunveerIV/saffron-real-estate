package util

import (
	"net/mail"
)

func ValidateEmail(email string) bool {
	addr, err := mail.ParseAddress(email)
	return err == nil && addr.Address == email
}
