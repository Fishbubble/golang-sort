package qsort

import (
// "fmt"
)

func Qsort(arr *[]int, l, r int) {
	if l < r {
		i := l
		j := r
		x := (*arr)[l]

		for {
			if i < j {
				for {
					if i < j && x <= (*arr)[j] {
						j--
					} else {
						break
					}
				}
				if i < j {
					(*arr)[i] = (*arr)[j]
					i++
				}

				for {
					if i < j && x >= (*arr)[i] {
						i++
					} else {
						break
					}
				}
				if i < j {
					(*arr)[j] = (*arr)[i]
					j--
				}
			} else {
				break
			}
		}
		(*arr)[i] = x
		Qsort(arr, l, i-1)
		Qsort(arr, i+1, r)
	}
}
