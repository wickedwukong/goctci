package chapter1

import "testing"

func TestHasUniqueChars(t *testing.T) {
	if (!HasUniqueChar("")) {
        t.Error("Expected true, got false")
	}

   if (!HasUniqueChar("ab")) {
        t.Error("Expected false, got true")
	}

	if (!HasUniqueChar("a")) {
        t.Error("Expected true, got false")
	}

   if (HasUniqueChar("aa")) {
        t.Error("Expected false, got true")
	}

   if (HasUniqueChar("aba")) {
        t.Error("Expected false, got true")
	}


}