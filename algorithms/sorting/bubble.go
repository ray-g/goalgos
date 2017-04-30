package sorting

func BubbleSort(v Sortable) {
	for i := 0; i < v.Len(); i++ {
		for j := 0; j < v.Len(); j++ {
			if v.Less(i, j) {
				v.Swap(i, j)
			}
		}
	}
}
