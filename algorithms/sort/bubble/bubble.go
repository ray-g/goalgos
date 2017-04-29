package bubble

import (
	"github.com/ray-g/goalgos/algorithms/sort"
)

func Sort(v sort.Sortable) {
	for i := 0; i < v.Len(); i++ {
		for j := 0; j < v.Len(); j++ {
			if v.Less(i, j) {
				v.Swap(i, j)
			}
		}
	}
}
