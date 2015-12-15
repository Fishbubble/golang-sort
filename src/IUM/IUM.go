// IUM means intersection,union and minus

package IUM

import (
// ""
)

func IUM(arr1 []int, arr2 []int) ([]int, []int, []int) {
	arr1_size := len(arr1)
	arr2_size := len(arr2)
	end := arr1_size
	swap := false

	var intersection, union, minus []int

	for i := 0; i < end; {
		swap = false
		for j := i; j < arr2_size; j++ {
			if arr1[i] == arr2[j] {
				arr2[i], arr2[j] = arr2[j], arr2[i]
				swap = true
				break
			}
		}
		if !swap {
			end--
			arr1[i], arr1[end] = arr1[end], arr1[i]
		} else {
			i++
		}
	}

	intersection = arr1[0:end]

	union = arr1
	union = append(union, arr2[end:]...)

	minus = arr1[end:]
	return intersection, union, minus
}
