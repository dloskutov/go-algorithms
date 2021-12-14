package findmedian

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// FindMedian - find median for unsorted array
func FindMedian(array []float64) float64 {
	length := len(array)
	isOdd := (length % 2) == 1

	if isOdd {
		return quickselect(length/2, array)
	}
	return 0.5 * (quickselect(length/2, array) + quickselect(length/2-1, array))
}

func quickselect(index int, array []float64) float64 {
	length := len(array)
	if length == 1 {
		return array[0]
	}

	leftArray := make([]float64, 0)
	rightArray := make([]float64, 0)
	pivotsArray := make([]float64, 0)

	pivotValue := array[r.Intn(length)]

	for _, value := range array {
		if pivotValue > value {
			leftArray = append(leftArray, value)
		} else if pivotValue < value {
			rightArray = append(rightArray, value)
		} else {
			pivotsArray = append(pivotsArray, value)
		}
	}

	if index < len(leftArray) {
		return quickselect(index, leftArray)
	}
	if index < len(leftArray)+len(pivotsArray) {
		return pivotsArray[0]
	}

	return quickselect(index-len(leftArray)-len(pivotsArray), rightArray)
}
