package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const MAX = 1000000

func generateRandomArray(length int) []int64 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int64, 0)
	for i := 0; i < length; i++ {
		arr = append(arr, random.Int63n(MAX))
	}
	return arr
}

func isSorted(array []int64) bool {
	for i := 0; i+1 < len(array); i++ {
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	arr := generateRandomArray(1000)

	err := Sort(arr)
	assert.Equal(t, nil, err)
	assert.True(t, isSorted(arr))
}
