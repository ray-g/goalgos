package sorting

import (
	"reflect"
	"sort"
	"testing"
)

type IntSlice []int

func (ns IntSlice) Len() int           { return len(ns) }
func (ns IntSlice) Less(i, j int) bool { return ns[i] < ns[j] }
func (ns IntSlice) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }

func testSort(t *testing.T, name string, foo func(s Sortable)) {
	tcs := [][]int{
		[]int{3, 1, 2},
	}

	for _, tc := range tcs {
		expects := make([]int, len(tc))
		actual := make([]int, len(tc))

		copy(expects, tc)
		copy(actual, tc)

		sort.Sort(IntSlice(expects))
		foo(IntSlice(actual))

		if !reflect.DeepEqual(actual, expects) {
			t.Errorf("Error \"%s\"\nInput: %v\nExpects: %v\nActual: %v", name, tc, expects, actual)
		}
	}
}

func TestBubbleSort(t *testing.T) { testSort(t, "BubbleSort", BubbleSort) }
