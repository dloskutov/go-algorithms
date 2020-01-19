package other

import "testing"

import "github.com/stretchr/testify/assert"

func TestIsPrime(t *testing.T) {
	assert.Equal(t, true, IsPrime(2))
	assert.Equal(t, true, IsPrime(3))
	assert.Equal(t, true, IsPrime(13))
	assert.Equal(t, true, IsPrime(17))
	assert.Equal(t, true, IsPrime(233))
	assert.Equal(t, true, IsPrime(997))

	assert.Equal(t, false, IsPrime(1))
	assert.Equal(t, false, IsPrime(4))
	assert.Equal(t, false, IsPrime(8))
	assert.Equal(t, false, IsPrime(81))
	assert.Equal(t, false, IsPrime(243))
	assert.Equal(t, false, IsPrime(998))
}
