package priorityqueue

import (
	"fmt"

	heap "github.com/ray-g/goalgos/data-structures/heaps/binary-heap"
	queue "github.com/ray-g/goalgos/data-structures/queue"
)

type Item struct {
	Value    interface{}
	Priority int
}

func NewItem(v interface{}, p int) (n *Item) {
	return &Item{
		Value:    v,
		Priority: p,
	}
}

func (n Item) Less(v interface{}) bool {
	return n.Priority < v.(Item).Priority
}

type PriorityQueue struct {
	data heap.Heap
}

func New(isMin bool) (q *PriorityQueue) {
	return &PriorityQueue{
		data: *heap.New(isMin),
	}
}

func (pq *PriorityQueue) Len() int {
	return pq.data.Len()
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.data.IsEmpty()
}

func (pq *PriorityQueue) Insert(n Item) {
	pq.data.Insert(n)
}

func (pq *PriorityQueue) Extract() (n Item) {
	return pq.data.Extract().(Item)
}

func (pq *PriorityQueue) Top() (n Item) {
	return pq.data.Top().(Item)
}

func (pq *PriorityQueue) ChangePriority(v interface{}, p int) (err error) {
	q := queue.New()

	pop := pq.Extract()
	for v != pop.Value {
		if pq.Len() == 0 {
			err = fmt.Errorf("Item of %v not found", v)
			return
		}

		q.EnQueue(pop)
		pop = pq.Extract()
	}

	pop.Priority = p
	pq.Insert(pop)

	for q.Size() > 0 {
		pq.Insert(q.DeQueue().(Item))
	}

	return
}
