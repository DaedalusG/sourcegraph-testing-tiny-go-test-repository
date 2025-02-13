package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	dog := Dog{}
	cat := Cat{}

	assert.NotEqual(t, dog.Sound(), cat.Sound(), "make different sound")
}
