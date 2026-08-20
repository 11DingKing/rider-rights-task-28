package domain

import "strings"

// ValidateModificationActor ensures the operator recorded against an item
// modification refers to an attributable identity. A value that is empty or
// made up entirely of whitespace cannot be traced back to anyone, so the
// audit trail would be unable to tell who changed the material; reject it.
func ValidateModificationActor(actor string) error {
	if strings.TrimSpace(actor) == "" {
		return ValidationError{Field: "actor", Message: "must not be empty or whitespace"}
	}
	return nil
}

// NormalizeModificationActor trims surrounding whitespace from the operator
// identity so the audit trail stores a clean, attributable identity. It is the
// companion to ValidateModificationActor: validate first to reject blank
// input, then normalize before persisting the audit entry.
func NormalizeModificationActor(actor string) string {
	return strings.TrimSpace(actor)
}
