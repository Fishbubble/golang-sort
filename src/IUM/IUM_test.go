package IUM

import (
	"testing"
)

func TestIUM(t *testing.T) {
	var arr1 []int = []int{5, 2, 4}
	var arr2 []int = []int{3, 6, 2, 1}
	i, u, m := IUM(arr1, arr2)
	if i[0] != 2 || len(i) != 1 {
		t.Error("Intersection error,exception is 2,got ", i[0])
	}
	if len(u) != 6 {
		t.Error("Union error,exception len is 6,got ", len(u))
	}
	if len(m) != 2 {
		t.Error("Minus error,exception len is 2,got ", len(m))
	}

	t.Error(i, u, m)
}

func TestIUM1(t *testing.T) {
	var arr1 []int = []int{5, 4, 3, 2, 1}
	var arr2 []int = []int{1, 2, 3, 4, 5}
	i, u, m := IUM(arr1, arr2)
	if i[0] != 5 || len(i) != 5 {
		t.Error("Intersection error,exception len is 5,got ", len(i))
	}
	if len(u) != 5 {
		t.Error("Union error,exception len is 5,got ", len(u))
	}
	if len(m) != 0 {
		t.Error("Minus error,exception len is 0,got ", len(m))
	}

	for j := 0; j < len(arr1); j++ {
		if i[j] != u[j] {
			t.Error("Result error,i should be equal u,but \ni = ", i, "\nu = ", u)
			break
		}
	}
	t.Error(i, u, m)
}

func TestIUM2(t *testing.T) {
	var arr1 []int = []int{5, 4, 3, 2, 1}
	var arr2 []int = []int{6, 7, 8, 9, 10}
	i, u, m := IUM(arr1, arr2)
	if len(i) != 0 {
		t.Error("Intersection error,exception len is 0,got ", len(i))
	}
	if len(u) != 10 {
		t.Error("Union error,exception len is 10,got ", len(u))
	}
	if len(m) != 5 {
		t.Error("Minus error,exception len is 5,got ", len(m))
	}

	for j := 0; j < len(arr1); j++ {
		if arr1[j] != m[j] {
			t.Error("Result error,m should be equal arr1,but \nm \t= ", m, "\narr1 \t= ", arr1)
			break
		}
	}
	t.Error(i, u, m)
}
