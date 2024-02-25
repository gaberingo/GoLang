package challenge_list

import (
	"fmt"
)

func MoveZeros(arr []int) []int {
	tempArr := []int{}
	for _, num := range arr {
		if num != 0 {
			tempArr = append(tempArr, num)
		}
	}
	zeroNeed := len(arr) - len(tempArr)
	for i := 0; i < zeroNeed; i++ {
		fmt.Println(i)
		fmt.Println(len(arr), len(tempArr))
		tempArr = append(tempArr, 0)
	}
	return tempArr
}
