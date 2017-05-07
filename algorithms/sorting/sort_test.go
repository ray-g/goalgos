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

func makeTestCases() [][]int {
	return [][]int{
		[]int{},
		[]int{3, 1, 2},
		[]int{3, 3, 1, 2, 2, 1},
		[]int{1, 3, 5, 4, 8, 7, 9, 6, 2, 0},
	}
}

func testSort(t *testing.T, name string, foo func(s Sortable)) {
	tcs := makeTestCases()

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

func TestBubbleSort(t *testing.T)    { testSort(t, "BubbleSort", BubbleSort) }
func TestQuickSort(t *testing.T)     { testSort(t, "QuickSort", QuickSort) }
func TestHeapSort(t *testing.T)      { testSort(t, "HeapSort", HeapSort) }
func TestSelectionSort(t *testing.T) { testSort(t, "SelectionSort", SelectionSort) }
func TestInsertionSort(t *testing.T) { testSort(t, "InsertionSort", InsertionSort) }
