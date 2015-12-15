package qsort

import (
	"testing"
)

func TestQsort(t *testing.T) {
	arr := []int{8, 3, 1, 6, 5, 7, 2, 4, 9, 0}
	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Qsort(&arr, 0, len(arr)-1)
	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			t.Error("Qsort error:\n\texpection\t= ", exp, "\n", "\tgot\t\t= ", arr)
			break
		}
	}
}

func TestQsort1(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}
	exp := []int{0, 1, 2, 3, 4, 5}
	Qsort(&arr, 0, len(arr)-1)
	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			t.Error("Qsort error:\n\texpection\t= ", exp, "\n", "\tgot\t\t= ", arr)
			break
		}
	}
}

func TestQsort2(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 0}
	exp := []int{0, 1, 2, 3, 4, 5}
	Qsort(&arr, 0, len(arr)-1)
	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			t.Error("Qsort error:\n\texpection\t= ", exp, "\n", "\tgot\t\t= ", arr)
			break
		}
	}
	t.Log("arr ", arr)
}
