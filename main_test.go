package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hello(t *testing.T) {
	got := Hello()

	assert.Equal(t, "Hello", got, "should work")
}
