package sorting

func InsertionSort(v Sortable) {
	for i := 1; i < v.Len(); i++ {
		for j := i; j > 0 && v.Less(j, j-1); j-- {
			v.Swap(j, j-1)
		}
	}
}
