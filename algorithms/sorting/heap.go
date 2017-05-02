package sorting

func sink(v Sortable, k, n int) {
	for 2*k <= n {
		j := 2 * k
		//if j < n && v.Less(j, j+1) {
		if j < n && v.Less(j-1, j) {
			j++
		}
		//if !v.Less(k, j) {
		if !v.Less(k-1, j-1) {
			break
		}
		//v.Swap(k, j)
		v.Swap(k-1, j-1)
		k = j
	}
}

func HeapSort(v Sortable) {
	n := v.Len()
	for k := n / 2; k >= 1; k-- {
		sink(v, k, n)
	}

	for n > 1 {
		// v.Swap(1, n)
		v.Swap(0, n-1)
		n--
		sink(v, 1, n)
	}
}
