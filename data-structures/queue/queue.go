package queue

import "sync"

type element struct {
	data interface{}
	next *element
}

type Queue struct {
	head *element
	tail *element
	lock *sync.Mutex
	size int
}

func New() *Queue {
	q := new(Queue)
	q.lock = &sync.Mutex{}

	return q
}

func (q *Queue) EnQueue(data interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	ele := new(element)
	ele.data = data

	if q.size == 0 {
		q.head, q.tail = ele, ele
	} else {
		q.tail.next, q.tail = ele, ele
	}

	q.size++
}

func (q *Queue) DeQueue() (data interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		return nil
	}

	data, q.head = q.head.data, q.head.next
	q.size--
	return
}

func (q *Queue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.size
}

func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.size == 0
}
