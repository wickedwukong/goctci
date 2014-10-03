package chapter1

import (
	"testing"
 	"github.com/stretchr/testify/assert"
 )

func TestReplaceSpacesInEmptyString(t *testing.T) {
	value := ""
	ReplaceSpaces(&value)
	assert.Equal(t, value, "", "It should still be an empty string")
}

func TestReplaceSpaceInStringWithNoSpaces(t *testing.T) {
	value := "a"
	ReplaceSpaces(&value)
	assert.Equal(t, value, "a", "Nothing should be replaced")
}

func TestSingleSpaceShouldBeTrimmed(t *testing.T) {
	value := " "
	ReplaceSpaces(&value)
	assert.Equal(t, value, "", "The single space should be trimmed.")
}

func TestReplaceSpacesInMixedString(t *testing.T) {
	value := "a b  c"
	ReplaceSpaces(&value)
	assert.Equal(t, value, "a%20b%20%20c", "Multiple spaces should be replaced with %20")
}

func TestSpacesAtBeginingAndEndAreTrimmed(t *testing.T) {
	value := " a b  c  "
	ReplaceSpaces(&value)
	assert.Equal(t, value, "a%20b%20%20c", "Multiple spaces should be replaced with %20")
}