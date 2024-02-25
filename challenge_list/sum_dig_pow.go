package challenge_list

import (
	"math"
	"strconv"
)

func SumDigPow(a, b uint64) []uint64 {
	var arr []uint64
	for i := a; i <= b; i++ {
		strNum := strconv.FormatUint(i, 10)
		var sum float64
		for j, num := range strNum {
			base := float64(uint64(num - '0'))
			exponent := float64(j + 1)
			sum += math.Pow(base, exponent)
		}
		if i == uint64(sum) {
			arr = append(arr, i)
		}
	}
	return arr
}
