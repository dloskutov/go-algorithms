package sorting

// ShellSort - shell sorting
func ShellSort(array []int64) {
	length := len(array)

	var step = 1
	for step < length/3 {
		step += step*3 + 1
	}
	for step >= 1 {
		for i := step; i < length; i++ {
			for j := i; j >= step && array[j-step] > array[j]; j -= step {
				array[j-step], array[j] = array[j], array[j-step]
			}
		}

		step = step / 3
	}
}
