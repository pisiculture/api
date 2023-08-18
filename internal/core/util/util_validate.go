package util

import (
	"errors"
	"strings"
)

func ValidateEmail(email string) error {

	if !strings.Contains(email, "@") {
		return errors.New("Email not with @")
	}

	if !strings.HasSuffix(email, ".com") {
		return errors.New("Email does not end with .com")
	}

	return nil
}
