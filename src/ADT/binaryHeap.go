// Package adt description
package adt

// CreateBinaryHeap - create binary heap
func CreateBinaryHeap(values []int64) *binaryHeap {
	heap := &binaryHeap{
		array: make([]int64, 0, len(values)),
	}

	for index := range values {
		heap.Insert(values[index])
	}

	return heap
}

type binaryHeap struct {
	array []int64
}

func (heap *binaryHeap) Insert(value int64) {
	heap.array = append(heap.array, value)
	size := len(heap.array)

	index := size - 1
	for index > 0 {
		parentIndex := index / 2
		if index%2 == 0 {
			parentIndex = index/2 - 1
		}
		if parentIndex >= 0 && heap.array[index] > heap.array[parentIndex] {
			heap.array[index], heap.array[parentIndex] = heap.array[parentIndex], heap.array[index]
			index = parentIndex
			continue
		}
		break
	}
}

func (heap *binaryHeap) IsEmpty() bool {
	return len(heap.array) == 0
}

func (heap *binaryHeap) Get() int64 {
	size := len(heap.array)
	if size == 0 {
		return 0
	}

	value := heap.array[0]
	heap.array[0] = heap.array[size-1]
	heap.array = heap.array[:size-1]
	newSize := len(heap.array)

	index := 0
	for index*2+1 < newSize && heap.array[index] < heap.array[index*2+1] {
		leftChildIndex := index*2 + 1
		rightChildIndex := index*2 + 2
		leftChildValue := heap.array[leftChildIndex]

		if index*2+2 < newSize {
			rightChildValue := heap.array[rightChildIndex]

			if leftChildValue > rightChildValue {
				heap.array[index], heap.array[leftChildIndex] = heap.array[leftChildIndex], heap.array[index]
				index = leftChildIndex
			} else {
				heap.array[index], heap.array[rightChildIndex] = heap.array[rightChildIndex], heap.array[index]
				index = rightChildIndex
			}
		} else {
			heap.array[index], heap.array[leftChildIndex] = heap.array[leftChildIndex], heap.array[index]
			index = leftChildIndex
		}
	}
	return value
}
