package queue

import "sync"

// QueueS is slice implemented Queue
type QueueS struct {
	queue []interface{}
	lock  *sync.Mutex
	size  int
}

func NewSQ() *QueueS {
	q := &QueueS{}
	q.queue = make([]interface{}, 0)
	q.lock = &sync.Mutex{}

	return q
}

func (q *QueueS) EnQueue(ele interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queue = append(q.queue, ele)

	q.size++
}

func (q *QueueS) DeQueue() (ele interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		return nil
	}

	ele, q.queue = q.queue[0], q.queue[1:]
	q.size--
	return
}

func (q *QueueS) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.size
}

func (q *QueueS) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.size == 0
}
