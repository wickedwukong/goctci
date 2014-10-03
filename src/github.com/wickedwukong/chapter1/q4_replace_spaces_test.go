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

func TestReplaceSingleSpace(t *testing.T) {
	value := " "
	ReplaceSpaces(&value)
	assert.Equal(t, value, "%20", "The single space should be replaced with %20")
}