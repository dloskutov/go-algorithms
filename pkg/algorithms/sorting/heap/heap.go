package heap

import "github.com/dloskutov/go-algorithms/pkg/datastructures/binaryheap"

func Sort(array []int64) error {
	binaryHeap, err := binaryheap.New(nil, binaryheap.Min)
	if err != nil {
		return err
	}
	for i := range array {
		err = binaryHeap.Insert(int(array[i]), array[i])
		if err != nil {
			return err
		}
	}

	index := 0
	for {
		value, err := binaryHeap.Top()
		if err == binaryheap.ErrBinaryHeapIsEmpty {
			return nil
		}
		if err != nil {
			return err
		}
		array[index] = value.(int64)
		index++
	}
}
