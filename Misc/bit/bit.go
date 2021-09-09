package bit

import "math"

func GetMin1(base, exp int) int {
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
			exp >>= 1
		}
		if exp != 1 {
			break
		}
		base *= base
	}

	return result
}

func GetMin2(in, in2 int) int {
	return int(math.Pow(float64(in), float64(in2)))
}
