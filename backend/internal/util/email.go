package util

import "errors"

func ValidateEmail(email string) error {
	if email != "" {
		return nil
	}
	return errors.New("invalid Email")
}
