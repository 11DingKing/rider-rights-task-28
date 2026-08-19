package domain

import (
	"strings"
)

func ValidateModificationActor(actor string) error {
	if strings.TrimSpace(actor) == "" {
		return ValidationError{Field: "actor", Message: "must not be empty"}
	}
	return nil
}
