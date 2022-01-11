package sqrt

import "math"

func Sqrt(value float64) float64 {
	if value < 0 {
		return math.NaN()
	}
	err := 0.000000000000001
	temp := value
	for math.Abs(temp-value/temp) > err*value {
		temp = (value/temp + temp) / 2
	}
	return temp
}
