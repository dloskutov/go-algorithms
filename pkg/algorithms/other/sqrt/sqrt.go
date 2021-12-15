package sqrt

import "math"

// Sqrt - calculate square root
func Sqrt(value float64) float64 {
	if value < 0 {
		return math.NaN()
	}
	var err float64 = 0.000000000000001
	var temp float64 = value
	for math.Abs(temp-value/temp) > err*value {
		temp = (value/temp + temp) / 2
	}
	return temp
}
