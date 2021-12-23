package aboutSort

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	i := 3
	ret := BinarySearch(0, 5, i, nums)
	t.Logf("ret = %v\n", ret)
}

func TestSelectionSort(t *testing.T) {
	nums := []int{1, 3,2,4,5,6,7,9,8,0, 4, 5, 6}
	ret:=SelectionSort(nums)
	t.Logf("after selectSort: %v",ret)
}