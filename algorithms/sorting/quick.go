package sorting

func QuickSort(v Sortable) {
	var partition func(left, right int)
	var sort func(left, right, pivot int) int

	sort = func(left, right, pivot int) int {
		right--
		v.Swap(pivot, right)
		for i := left; i < right; i++ {
			if v.Less(i, right) {
				v.Swap(i, left)
				left++
			}
		}
		v.Swap(left, right)
		return left
	}

	partition = func(left, right int) {
		if left < right {
			pivot := (left + right) / 2
			pivot = sort(left, right, pivot)
			partition(left, pivot)
			partition(pivot+1, right)
		}
	}

	partition(0, v.Len())
}
