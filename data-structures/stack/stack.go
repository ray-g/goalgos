package stack

import "sync"

type element struct {
	data interface{}
	next *element
}

type Stack struct {
	top  *element
	lock *sync.Mutex
	size int
}

func New() *Stack {
	s := new(Stack)
	s.lock = &sync.Mutex{}

	return s
}

func (s *Stack) Push(data interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	ele := new(element)
	ele.data = data
	ele.next = s.top
	s.top = ele

	s.size++
}

func (s *Stack) Pop() (data interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top == nil {
		return nil
	}

	data, s.top = s.top.data, s.top.next
	s.size--
	return
}

func (s *Stack) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.size
}

func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.size == 0
}
