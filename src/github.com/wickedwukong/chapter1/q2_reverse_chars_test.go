package chapter1

import (
	"testing"
 	"github.com/stretchr/testify/assert"
 )

func TestReverseEmptyString(t *testing.T) {
	value := ""
	Reverse(&value)
	assert.Equal(t, value, "", "'' should be reversed to be ''")
}

func TestReverseSingleCharString(t *testing.T) {
	value := "a"
	Reverse(&value)
	assert.Equal(t, value, "a", "'a' should be reversed to be 'a'")
}

func TestReverseTwoRepeatedCharString(t *testing.T) {
	value := "aa"
	Reverse(&value)
	assert.Equal(t, value, "aa", "'aa' should be reversed to be 'aa'")
}

func TestReverseLongString(t *testing.T) {
	value := "abcd"
	Reverse(&value)
	assert.Equal(t, value, "dcba", "'abcd' should be reversed to be 'dcba'")
}

func TestReverseNilStringPointer(t *testing.T) {
	var value *string
	Reverse(value)
	assert.Equal(t, value, (*string)(nil), "nil string pointer should be still be a nil string pointer")
}