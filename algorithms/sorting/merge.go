package sorting

import number "github.com/ray-g/goalgos/math/number-theory"

func merge(a, aux []int, lo, mid, hi int) {
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func MergeSortDown(a []int) {
	aux := make([]int, len(a))
	var sort func(a []int, lo, hi int)

	sort = func(a []int, lo, hi int) {
		if hi <= lo {
			return
		}
		mid := lo + (hi-lo)/2
		sort(a, lo, mid)
		sort(a, mid+1, hi)
		merge(a, aux, lo, mid, hi)
	}

	sort(a, 0, len(a)-1)
}

func MergeSortUp(a []int) {
	n := len(a)
	aux := make([]int, n)
	for sz := 1; sz < n; sz = sz + sz {
		for lo := 0; lo < n-sz; lo += sz + sz {
			merge(a, aux, lo, lo+sz-1, number.Min(lo+sz+sz-1, n-1))
		}
	}
}
