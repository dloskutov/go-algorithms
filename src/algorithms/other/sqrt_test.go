package other

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	assert.Equal(t, 1.414213562373095, Sqrt(2))
	assert.Equal(t, 3.0, Sqrt(9))
	assert.Equal(t, 5.0, Sqrt(25))

	assert.Equal(t, true, math.IsNaN(Sqrt(-1)))
}
