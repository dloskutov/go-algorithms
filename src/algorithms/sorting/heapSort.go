package sorting

import "github.com/dloskutov/go-algorithms/src/ADT/adt"

// HeapSort - heap sorting
func HeapSort(array []int64) {
	size := len(array)
	heap := adt.CreateBinaryHeap(array)

	index := 0
	for index < size {
		array[index] = heap.Get()
		index++
	}
}
