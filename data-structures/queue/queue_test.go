package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()

	if !q.IsEmpty() ||
		q.size != 0 ||
		q.Size() != 0 {
		t.Error()
	}

	q.EnQueue(0)
	q.EnQueue(1)
	q.EnQueue(2)

	if q.Size() != 3 {
		t.Error()
	}

	e0 := q.DeQueue()

	if e0 != 0 || q.Size() != 2 {
		t.Error()
	}

	e1 := q.DeQueue()

	if e1 != 1 {
		t.Error()
	}

	e2 := q.DeQueue()

	if e2 != 2 || q.Size() != 0 {
		t.Error()
	}

	ee := q.DeQueue()

	if ee != nil || q.Size() != 0 {
		t.Error()
	}
}

func TestQueueS(t *testing.T) {
	q := NewSQ()

	if !q.IsEmpty() ||
		q.size != 0 ||
		q.Size() != 0 {
		t.Error()
	}

	q.EnQueue(0)
	q.EnQueue(1)
	q.EnQueue(2)

	if q.Size() != 3 {
		t.Error()
	}

	e0 := q.DeQueue()

	if e0 != 0 || q.Size() != 2 {
		t.Error()
	}

	e1 := q.DeQueue()

	if e1 != 1 {
		t.Error()
	}

	e2 := q.DeQueue()

	if e2 != 2 || q.Size() != 0 {
		t.Error()
	}

	ee := q.DeQueue()

	if ee != nil || q.Size() != 0 {
		t.Error()
	}
}
