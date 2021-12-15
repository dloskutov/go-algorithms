package findmedian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedian(t *testing.T) {
	assert.Equal(t, 5.0, FindMedian([]float64{1, 2, 9, 4, 5, 3, 7, 8, 6}))
	assert.Equal(t, 5.5, FindMedian([]float64{7, 3, 5, 6, 8, 9, 4, 1, 2, 6}))
	assert.Equal(t, 2.0, FindMedian([]float64{5, 2, 1, 4, 2}))
	assert.Equal(t, 4.0, FindMedian([]float64{1, 5, 5, 1, 4}))
	assert.Equal(t, 100.0, FindMedian([]float64{1, 1000, 100}))
	assert.Equal(t, 51.0, FindMedian([]float64{1, 1000, 100, 2}))
	assert.Equal(t, 5.0, FindMedian([]float64{9, 1, 0, 2, 3, 4, 6, 8, 7, 10, 5}))
}
