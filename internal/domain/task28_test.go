package domain

import "testing"

func TestModificationRequiresActor(t *testing.T) {
	if err := ValidateModificationActor(" \t "); err == nil {
		t.Fatal("blank modification actor was accepted")
	}
	if err := ValidateModificationActor("prosecutor-1"); err != nil {
		t.Fatalf("valid actor rejected: %v", err)
	}
}
