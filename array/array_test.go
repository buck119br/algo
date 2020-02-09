package array

import "testing"

func TestSortedArrayMerge(t *testing.T) {
	a := []interface{}{1, 3, 3, 5, 7, 8, 9}
	b := []interface{}{2, 4, 5, 6, 7, 8, 10}
	t.Log(SortedArrayMerge(a, b, func(x, y interface{}) bool { return x.(int) < y.(int) }))
}
