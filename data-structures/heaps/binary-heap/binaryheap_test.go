package heap

import "testing"

type Int int

func (x Int) Less(than interface{}) bool {
	return x < than.(Int)
}

func TestMinHeap(t *testing.T) {
	h := New(true)

	nums := []int{8, 7, 6, 3, 1, 0, 0, 0, 9, 0, 2, 4, 9, 5}
	for _, n := range nums {
		h.Insert(Int(n))
	}

	if 0 != h.Top().(Int) {
		t.Error()
	}

	ordered := make([]Int, 0)
	for h.Len() > 0 {
		ordered = append(ordered, h.Extract().(Int))
	}

	if nil != h.Extract() {
		t.Error()
	}

	if nil != h.Top() {
		t.Error()
	}

	for i := 0; i < len(ordered)-2; i++ {
		if ordered[i] > ordered[i+1] {
			t.Errorf("!Ordered. %v", ordered)
		}
	}
}

func TestMaxHeap(t *testing.T) {
	h := New(false)

	nums := []int{8, 7, 6, 3, 1, 9, 9, 9, 9, 0, 2, 4, 9, 5}
	for _, n := range nums {
		h.Insert(Int(n))
	}

	if 9 != h.Top().(Int) {
		t.Error()
	}

	ordered := make([]Int, 0)
	for h.Len() > 0 {
		ordered = append(ordered, h.Extract().(Int))
	}

	if nil != h.Extract() {
		t.Error()
	}

	if nil != h.Top() {
		t.Error()
	}

	for i := 0; i < len(ordered)-2; i++ {
		if ordered[i] < ordered[i+1] {
			t.Errorf("!Ordered. %v", ordered)
		}
	}
}
