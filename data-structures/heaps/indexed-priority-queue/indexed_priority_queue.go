package indexedpriorityqueue

import (
	"fmt"
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

type IndexPriorityQueue struct {
	pq    []int // 1 based priority queue
	qp    []int // inverse of pq - qp[pq[i]] = pq[qp[i]] = i
	data  []interface{}
	cap   int
	len   int
	minPQ bool
}

func New(isMin bool, max int) (q *IndexPriorityQueue) {
	if max <= 0 {
		return nil
	}
	q = &IndexPriorityQueue{
		pq:    make([]int, max+1),
		qp:    make([]int, max+1),
		data:  make([]interface{}, max+1),
		len:   0,
		cap:   max,
		minPQ: isMin,
	}

	for i := 0; i <= max; i++ {
		q.qp[i] = -1
	}
	return
}

func (ipq *IndexPriorityQueue) Len() int {
	return ipq.len
}

func (ipq *IndexPriorityQueue) IsEmpty() bool {
	return ipq.len == 0
}

func (ipq *IndexPriorityQueue) Contains(i int) bool {
	if i <= 0 || i >= ipq.cap {
		return false
	}

	return ipq.qp[i] != -1
}

func (ipq *IndexPriorityQueue) Insert(i int, n Item) (err error) {
	if i <= 0 || i >= ipq.cap {
		err = fmt.Errorf("Index out of range")
		return
	}

	if ipq.Contains(i) {
		err = fmt.Errorf("Index is already in the priority queue")
		return
	}
	ipq.len++
	ipq.qp[i] = ipq.len
	ipq.pq[ipq.len] = i
	ipq.data[i] = n
	ipq.swim(ipq.len)
	return
}

func (ipq *IndexPriorityQueue) TopIndex() int {
	if ipq.len == 0 {
		return -1
	}

	return ipq.pq[1]
}

func (ipq *IndexPriorityQueue) Top() interface{} {
	if ipq.len == 0 {
		return nil
	}

	return ipq.data[ipq.pq[1]]
}

func (ipq *IndexPriorityQueue) Extract() interface{} {
	if ipq.len == 0 {
		return nil
	}

	top := ipq.pq[1]
	ipq.swap(1, ipq.len)
	ipq.len--
	ipq.sink(1)
	ipq.qp[top] = -1
	n := ipq.data[top]
	ipq.data[top] = nil
	ipq.pq[ipq.len+1] = -1
	return n
}

func (ipq *IndexPriorityQueue) At(i int) (n interface{}, err error) {
	if i <= 0 || i >= ipq.cap {
		err = fmt.Errorf("Index out of range")
		return
	}

	if !ipq.Contains(i) {
		err = fmt.Errorf("Index not exists")
		return
	}

	n = ipq.data[i].(Item)
	return
}

func (ipq *IndexPriorityQueue) ChangeValue(i int, n Item) (err error) {
	if i <= 0 || i >= ipq.cap {
		err = fmt.Errorf("Index out of range")
		return
	}
	if !ipq.Contains(i) {
		err = fmt.Errorf("Index not exists")
		return
	}
	ipq.data[i] = n
	ipq.swim(ipq.qp[i])
	ipq.sink(ipq.qp[i])
	return
}

func (ipq *IndexPriorityQueue) DecreasePriority(i, prior int) (err error) {
	if i <= 0 || i >= ipq.cap {
		err = fmt.Errorf("Index out of range")
		return
	}
	if !ipq.Contains(i) {
		err = fmt.Errorf("Index not exists")
		return
	}
	if ipq.data[i].(Item).Priority <= prior {
		err = fmt.Errorf("Priority not changeable")
		return
	}

	ipq.data[i] = Item{ipq.data[i].(Item).Value, prior}
	if ipq.minPQ {
		ipq.swim(ipq.qp[i])
	} else {
		ipq.sink(ipq.qp[i])
	}
	return
}

func (ipq *IndexPriorityQueue) IncreasePriority(i, prior int) (err error) {
	if i <= 0 || i >= ipq.cap {
		err = fmt.Errorf("Index out of range")
		return
	}
	if !ipq.Contains(i) {
		err = fmt.Errorf("Index not exists")
		return
	}
	if ipq.data[i].(Item).Priority >= prior {
		err = fmt.Errorf("Priority not changeable")
		return
	}

	ipq.data[i] = Item{ipq.data[i].(Item).Value, prior}
	if ipq.minPQ {
		ipq.sink(ipq.qp[i])
	} else {
		ipq.swim(ipq.qp[i])
	}
	return
}

func (ipq *IndexPriorityQueue) Delete(i int) (err error) {
	if i <= 0 || i >= ipq.cap {
		err = fmt.Errorf("Index out of range")
		return
	}
	if !ipq.Contains(i) {
		err = fmt.Errorf("Index not exists")
		return
	}
	index := ipq.qp[i]
	ipq.swap(index, ipq.len)
	ipq.len--
	ipq.swim(index)
	ipq.sink(index)
	ipq.data[i] = nil
	ipq.qp[i] = -1
	return
}

func (ipq *IndexPriorityQueue) less(i, j int) bool {
	if ipq.minPQ {
		return ipq.data[ipq.pq[i]].(Item).Less(ipq.data[ipq.pq[j]])
	}
	return ipq.data[ipq.pq[j]].(Item).Less(ipq.data[ipq.pq[i]])
}

func (ipq *IndexPriorityQueue) swap(i, j int) {
	ipq.pq[i], ipq.pq[j] = ipq.pq[j], ipq.pq[i]
	ipq.qp[ipq.pq[i]] = i
	ipq.qp[ipq.pq[j]] = j
}

func (ipq *IndexPriorityQueue) swim(k int) {
	for k > 1 && ipq.less(k, k/2) {
		ipq.swap(k, k/2)
		k = k / 2
	}
}

func (ipq *IndexPriorityQueue) sink(k int) {
	for 2*k <= ipq.len {
		j := 2 * k
		if j < ipq.len && ipq.less(j+1, j) {
			j++
		}
		if ipq.less(k, j) {
			break
		}
		ipq.swap(k, j)
		k = j
	}
}
