package merge

func Sort(array []int64) {
	arrayCopy := make([]int64, 0)
	arrayCopy = append(arrayCopy, array...)

	sort(array, arrayCopy, 0, len(array)-1)
}

func sort(array []int64, arrayCopy []int64, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	if startIndex+1 == endIndex && array[startIndex] > array[endIndex] {
		array[startIndex], array[endIndex] = array[endIndex], array[startIndex]
	}
	mid := startIndex + (endIndex-startIndex)/2

	sort(array, arrayCopy, startIndex, mid)
	sort(array, arrayCopy, mid+1, endIndex)
	merge(array, arrayCopy, startIndex, mid+1, endIndex)
}

func merge(array []int64, arrayCopy []int64, startIndex int, mid int, endIndex int) {
	for i := startIndex; i <= endIndex; i++ {
		arrayCopy[i] = array[i]
	}

	var startLeftArrIndex int = startIndex
	var startRightArrIndex int = mid
	for i := startIndex; i <= endIndex; i++ {
		if startLeftArrIndex >= mid {
			array[i] = arrayCopy[startRightArrIndex]
			startRightArrIndex++
		} else if startRightArrIndex > endIndex {
			array[i] = arrayCopy[startLeftArrIndex]
			startLeftArrIndex++
		} else if arrayCopy[startLeftArrIndex] > arrayCopy[startRightArrIndex] {
			array[i] = arrayCopy[startRightArrIndex]
			startRightArrIndex++
		} else {
			array[i] = arrayCopy[startLeftArrIndex]
			startLeftArrIndex++
		}
	}
}
