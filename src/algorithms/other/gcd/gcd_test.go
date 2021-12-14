package gcd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	assert.Equal(t, 2, GCD(6, 4))
	assert.Equal(t, 3, GCD(21, 9))
	assert.Equal(t, 7, GCD(49, 7))
	assert.Equal(t, 2, GCD(4, 2))

	assert.Equal(t, 1, GCD(6, 1))
	assert.Equal(t, 1, GCD(7, 3))
}
