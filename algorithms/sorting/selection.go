package sorting

func SelectionSort(v Sortable) {
	for i := 0; i < v.Len(); i++ {
		min := i
		for j := i + 1; j < v.Len(); j++ {
			if v.Less(j, min) {
				min = j
			}
		}
		v.Swap(i, min)
	}
}
