package sorting

func ShellSort(v Sortable) {
	h := 1
	for h < v.Len()/3 {
		h = 3*h + 1 // 1, 4, 13, 40, 121, 364, 1093, ...
	}
	for h > 0 {
		for i := h; i < v.Len(); i++ {
			for j := i; j >= h && v.Less(j, j-h); j -= h {
				v.Swap(j, j-h)
			}
		}
		h = h / 3
	}
}
