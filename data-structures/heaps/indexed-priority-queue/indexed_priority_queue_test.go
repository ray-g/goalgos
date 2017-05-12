package indexedpriorityqueue

import (
	"fmt"
	"testing"
)

func preparePriorityQueue(pq *IndexPriorityQueue) {
	pq.Insert(8, *NewItem(8, 8))
	pq.Insert(1, *NewItem(7, 7))
	pq.Insert(2, *NewItem(5, 5))
	pq.Insert(9, *NewItem(3, 3))
	pq.Insert(4, *NewItem(2, 2))
	pq.Insert(5, *NewItem(1, 1))
	pq.Insert(6, *NewItem(6, 6))
	pq.Insert(7, *NewItem(0, 0))
	pq.Insert(10, *NewItem(4, 4))
	pq.Insert(3, *NewItem(9, 9))
}

func TestMaxPriorityQueue(t *testing.T) {
	nilPQ := New(false, 0)
	if nilPQ != nil {
		t.Error()
	}

	maxPQ := New(false, 20)

	if true != maxPQ.IsEmpty() {
		t.Error()
	}

	preparePriorityQueue(maxPQ)

	if false != maxPQ.IsEmpty() {
		t.Error()
	}

	if 9 != maxPQ.Top().(Item).Value {
		t.Error()
	}

	if 3 != maxPQ.TopIndex() {
		t.Error()
	}

	sorted := make([]Item, 0)
	for maxPQ.Len() > 0 {
		sorted = append(sorted, maxPQ.Extract().(Item))
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i].Priority < sorted[i+1].Priority {
			fmt.Println(sorted)
			t.Error()
		}
	}
}

func TestMinPriorityQueue(t *testing.T) {
	nilPQ := New(false, -1)
	if nilPQ != nil {
		t.Error()
	}
	minPQ := New(true, 20)

	if true != minPQ.IsEmpty() {
		t.Error()
	}

	preparePriorityQueue(minPQ)

	if false != minPQ.IsEmpty() {
		t.Error()
	}

	if 0 != minPQ.Top().(Item).Value {
		t.Error()
	}

	if 7 != minPQ.TopIndex() {
		t.Error()
	}

	sorted := make([]Item, 0)
	for minPQ.Len() > 0 {
		sorted = append(sorted, minPQ.Extract().(Item))
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i].Priority > sorted[i+1].Priority {
			fmt.Println(sorted)
			t.Error()
		}
	}

	if minPQ.Top() != nil {
		t.Error()
	}

	if minPQ.Extract() != nil {
		t.Error()
	}

	if minPQ.TopIndex() != -1 {
		t.Error()
	}
}

func TestContains(t *testing.T) {
	tcs := []struct {
		index   int
		expects bool
	}{
		{0, false},
		{1, true},
		{15, false},
		{20, false},
		{50, false},
		{-1, false},
	}

	q := New(false, 20)
	preparePriorityQueue(q)

	for _, tc := range tcs {
		if tc.expects != q.Contains(tc.index) {
			t.Errorf("Contains index %d not equal to %v.", tc.index, tc.expects)
		}
	}
}

func TestInsertErrors(t *testing.T) {
	q := New(false, 20)
	preparePriorityQueue(q)

	for _, index := range []int{-1, 0, 21, 50} {
		err := q.Insert(index, Item{20, 20})
		if err.Error() != "Index out of range" {
			t.Error()
		}
	}

	for _, index := range []int{1, 2, 3, 5, 10} {
		err := q.Insert(index, Item{20, 20})
		if err.Error() != "Index is already in the priority queue" {
			t.Error()
		}
	}
}

func TestAt(t *testing.T) {
	q := New(false, 20)
	preparePriorityQueue(q)

	n, err := q.At(-1)
	if err.Error() != "Index out of range" {
		t.Error()
	}
	if n != nil {
		t.Error()
	}
	n, err = q.At(0)
	if err.Error() != "Index out of range" {
		t.Error()
	}
	if n != nil {
		t.Error()
	}
	n, err = q.At(50)
	if err.Error() != "Index out of range" {
		t.Error()
	}
	if n != nil {
		t.Error()
	}

	n, err = q.At(15)
	if err.Error() != "Index not exists" {
		t.Error()
	}
	if n != nil {
		t.Error()
	}
	for i := 1; i <= 10; i++ {
		n, err = q.At(i)
		if err != nil {
			t.Error()
		}
		if n == nil {
			t.Error()
		}
	}
}

func TestChangeValue(t *testing.T) {
	q := New(false, 20)
	preparePriorityQueue(q)

	n := Item{20, 20}
	err := q.ChangeValue(-1, n)
	if err.Error() != "Index out of range" {
		t.Error()
	}
	err = q.ChangeValue(0, n)
	if err.Error() != "Index out of range" {
		t.Error()
	}
	err = q.ChangeValue(50, n)
	if err.Error() != "Index out of range" {
		t.Error()
	}

	err = q.ChangeValue(15, n)
	if err.Error() != "Index not exists" {
		t.Error()
	}

	err = q.ChangeValue(1, n)
	if err != nil {
		t.Error()
	}

	if q.Top().(Item).Value != 20 {
		t.Error()
	}
}

func TestDelete(t *testing.T) {
	q := New(false, 20)
	preparePriorityQueue(q)

	err := q.Delete(-1)
	if err.Error() != "Index out of range" {
		t.Error()
	}
	err = q.Delete(0)
	if err.Error() != "Index out of range" {
		t.Error()
	}
	err = q.Delete(50)
	if err.Error() != "Index out of range" {
		t.Error()
	}

	err = q.Delete(15)
	if err.Error() != "Index not exists" {
		t.Error()
	}

	for i := 1; i <= 10; i++ {
		err = q.Delete(i)
		if err != nil {
			t.Error()
		}
	}

	for i := 1; i <= 10; i++ {
		err = q.Delete(i)
		if err.Error() != "Index not exists" {
			t.Error()
		}
	}
}

func TestIncreasePriority(t *testing.T) {
	maxPQ := New(false, 20)

	preparePriorityQueue(maxPQ)

	err := maxPQ.IncreasePriority(8, 88)
	if 8 != maxPQ.Top().(Item).Value || err != nil {
		t.Error()
	}

	err = maxPQ.IncreasePriority(12, 10)
	if err.Error() != "Index not exists" {
		t.Error()
	}

	err = maxPQ.IncreasePriority(0, 10)
	if err.Error() != "Index out of range" {
		t.Error()
	}

	err = maxPQ.IncreasePriority(-1, 10)
	if err.Error() != "Index out of range" {
		t.Error()
	}

	err = maxPQ.IncreasePriority(50, 10)
	if err.Error() != "Index out of range" {
		t.Error()
	}

	err = maxPQ.IncreasePriority(8, 6)
	if err.Error() != "Priority not changeable" {
		t.Error()
	}

	minPQ := New(true, 20)
	preparePriorityQueue(minPQ)
	minPQ.IncreasePriority(7, 50)
	if minPQ.Top().(Item).Value != 1 {
		t.Error()
	}
}

func TestDecreasePriority(t *testing.T) {
	maxPQ := New(false, 20)

	preparePriorityQueue(maxPQ)

	err := maxPQ.DecreasePriority(3, 2)
	if 8 != maxPQ.Top().(Item).Value || err != nil {
		t.Error()
	}

	err = maxPQ.DecreasePriority(12, 10)
	if err.Error() != "Index not exists" {
		t.Error()
	}

	err = maxPQ.DecreasePriority(0, 10)
	if err.Error() != "Index out of range" {
		t.Error()
	}

	err = maxPQ.DecreasePriority(-1, 10)
	if err.Error() != "Index out of range" {
		t.Error()
	}

	err = maxPQ.DecreasePriority(50, 10)
	if err.Error() != "Index out of range" {
		t.Error()
	}

	err = maxPQ.DecreasePriority(8, 9)
	if err.Error() != "Priority not changeable" {
		t.Error()
	}

	minPQ := New(true, 20)
	preparePriorityQueue(minPQ)
	minPQ.DecreasePriority(8, -1)
	if minPQ.Top().(Item).Value != 8 {
		t.Error()
	}
}
