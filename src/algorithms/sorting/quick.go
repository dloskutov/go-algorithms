package sorting

// Quick - quick sort in situ
func Quick(array []int64) {
	quick(array, 0, len(array)-1)
}

func quick(array []int64, indexStart int, indexEnd int) {
	if indexEnd <= indexStart {
		return
	}

	middle := partition(array, indexStart, indexEnd)
	quick(array, indexStart, middle-1)
	quick(array, middle+1, indexEnd)
}

func partition(array []int64, indexStart, indexEnd int) int {
	pivotIndex := indexEnd
	pivotValue := array[pivotIndex]
	leftIndex := indexStart
	rightIndex := indexStart

	for index := indexStart; index <= indexEnd; index++ {
		if index == pivotIndex {
			array[leftIndex], array[pivotIndex] = array[pivotIndex], array[leftIndex]
			break
		}
		value := array[index]
		if value < pivotValue {
			array[leftIndex], array[rightIndex] = array[rightIndex], array[leftIndex]
			leftIndex++
			rightIndex++
			continue
		}
		if value >= pivotValue {
			rightIndex++
			continue
		}
	}

	return leftIndex
}
