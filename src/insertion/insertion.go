package insertion

func InsertionSort(arr *[]int, l, r int) {
	var i, j int
	for i = l + 1; i <= r; i += 1 {
		temp := (*arr)[i]
		for j = i - 1; j >= l && (*arr)[j] > temp; j-- {
			(*arr)[j+1] = (*arr)[j]
		}
		(*arr)[j+1] = temp
	}
}
