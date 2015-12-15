package mix

import (
	"insertion"
	// "fmt"
)

// mix qsort and insertion sort
func Sort(arr *[]int, l, r int) {
	if r-l+1 <= 8 {
		// fmt.Println("1")
		insertion.InsertionSort(arr, l, r)
	} else {
		// fmt.Println("2")
		qsort(arr, l, r)
	}
}

func qsort(arr *[]int, l, r int) {
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

		if i-l <= 8 {
			// fmt.Println("insertion")
			insertion.InsertionSort(arr, l, i-1)
		} else {
			// fmt.Println("qsort")
			qsort(arr, l, i-1)
		}

		if r-i <= 8 {
			// fmt.Println("insertion")
			insertion.InsertionSort(arr, i+1, r)
		} else {
			// fmt.Println("qsort")
			qsort(arr, i+1, r)
		}
	}
}
