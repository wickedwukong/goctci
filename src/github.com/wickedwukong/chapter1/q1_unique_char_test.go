package chapter1

import (
	"testing"
 	"github.com/stretchr/testify/assert"
 )

func TestHasUniqueChars(t *testing.T) {
	assert.True(t, HasUniqueChar(""), "It should be true that empty string has unique chars")
	assert.True(t, HasUniqueChar("ab"), "It should be true [ab] has unique chars")
	assert.True(t, HasUniqueChar("a"), "It should be true [a] has unique chars")
	assert.False(t, HasUniqueChar("aa"), "It should be false [aa] has unique chars")
	assert.False(t, HasUniqueChar("aba"), "It should be false [aba] has unique chars")
	assert.False(t, HasUniqueChar("bacad"), "It should be false [aba] has unique chars")
}