package sorting

import (
	"github.com/dloskutov/go-algorithms/src/adt"
)

// HeapSort - heap sorting
func HeapSort(array []int64) {
	size := len(array)
	heap := adt.CreateBinaryHeap(array)

	index := size - 1
	for index >= 0 {
		array[index] = heap.Get()
		index--
	}
}
