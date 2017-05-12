package priorityqueue

import (
	"fmt"
	"testing"
)

func preparePriorityQueue(pq *PriorityQueue) {
	pq.Insert(*NewItem(8, 8))
	pq.Insert(*NewItem(7, 7))
	pq.Insert(*NewItem(5, 5))
	pq.Insert(*NewItem(3, 3))
	pq.Insert(*NewItem(2, 2))
	pq.Insert(*NewItem(1, 1))
	pq.Insert(*NewItem(6, 6))
	pq.Insert(*NewItem(0, 0))
	pq.Insert(*NewItem(4, 4))
	pq.Insert(*NewItem(9, 9))
}

func TestMaxPriorityQueue(t *testing.T) {
	maxPQ := New(false)

	if true != maxPQ.IsEmpty() {
		t.Error()
	}

	preparePriorityQueue(maxPQ)

	if false != maxPQ.IsEmpty() {
		t.Error()
	}

	if 9 != maxPQ.Top().Value {
		t.Error()
	}

	sorted := make([]Item, 0)
	for maxPQ.Len() > 0 {
		sorted = append(sorted, maxPQ.Extract())
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i].Priority < sorted[i+1].Priority {
			fmt.Println(sorted)
			t.Error()
		}
	}
}

func TestMinPriorityQueue(t *testing.T) {
	minPQ := New(true)

	if true != minPQ.IsEmpty() {
		t.Error()
	}

	preparePriorityQueue(minPQ)

	if false != minPQ.IsEmpty() {
		t.Error()
	}

	if 0 != minPQ.Top().Value {
		t.Error()
	}

	sorted := make([]Item, 0)
	for minPQ.Len() > 0 {
		sorted = append(sorted, minPQ.Extract())
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i].Priority > sorted[i+1].Priority {
			fmt.Println(sorted)
			t.Error()
		}
	}
}

func TestChangePriority(t *testing.T) {
	maxPQ := New(false)

	preparePriorityQueue(maxPQ)

	err := maxPQ.ChangePriority(8, 88)
	if 8 != maxPQ.Top().Value || err != nil {
		t.Error()
	}

	err = maxPQ.ChangePriority(10, 10)
	if err == nil {
		t.Error()
	}
}
