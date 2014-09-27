package chapter1

import "testing"

func TestHasUniqueChars(t *testing.T) {
	if (!HasUniqueChar("")) {
        t.Error("Expected 1.5, got ")
	}
}