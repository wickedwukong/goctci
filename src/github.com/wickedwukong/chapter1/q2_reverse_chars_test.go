package chapter1

import (
	"testing"
 	"github.com/stretchr/testify/assert"
 )

func TestReverse(t *testing.T) {
	value := "abcd"
	Reverse(&value)
	assert.Equal(t, value, "dcba", "'abcd' should be reversed to 'dcba'")
}