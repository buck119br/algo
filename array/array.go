package array

func SortedArrayMerge(a, b []interface{}, less func(x, y interface{}) bool) []interface{} {
	r := make([]interface{}, len(a)+len(b), len(a)+len(b))

	var i, j, x int
	for ; i < len(a) && j < len(b); x++ {
		if less(a[i], b[j]) {
			r[x] = a[i]
			i++
		} else {
			r[x] = b[j]
			j++
		}
	}

	for ; i < len(a); i++ {
		r[x] = a[i]
	}
	for ; j < len(b); j++ {
		r[x] = b[j]
	}

	return r
}
