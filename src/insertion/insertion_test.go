package insertion

import (
	"testing"
)

func TestInsertion(t *testing.T) {
	arr := []int{3, 2, 1}
	InsertionSort(&arr, 0, len(arr)-1)
	if len(arr) != 3 || arr[0] != 1 || arr[1] != 2 || arr[2] != 3 {
		t.Errorf("TestInsertion error arr[0]=%d arr[1]=%d arr[2]=%d", arr[0], arr[1], arr[2])
	}
	t.Logf("arr[0]=%d arr[1]=%d arr[2]=%d", arr[0], arr[1], arr[2])
}

func TestInsertion1(t *testing.T) {
	arr := []int{3, 2, 1, 1, 2, 3}
	InsertionSort(&arr, 0, len(arr)-1)
	if len(arr) != 6 || arr[0] != 1 || arr[1] != 1 || arr[2] != 2 {
		t.Errorf("TestInsertion error arr[0]=%d arr[1]=%d arr[2]=%d", arr[0], arr[1], arr[2])
	}
	t.Logf("arr[0]=%d arr[1]=%d arr[2]=%d arr[3]=%d arr[4]=%d arr[5]=%d", arr[0], arr[1], arr[2], arr[3], arr[4], arr[5])
}
