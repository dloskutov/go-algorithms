package heap

// @TODO: need to implement
func Sort(array []int64) {
	length := len(array)

	for i := 0; i < length; i++ {
		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}
	}
}
