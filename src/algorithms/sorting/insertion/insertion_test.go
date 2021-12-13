package sorting

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
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

func TestInsertion(t *testing.T) {
	arr := generateRandomArray(1000)

	Insertion(arr)
	assert.True(t, isSorted(arr))
}
