package stack

import "sync"

type element struct {
	data interface{}
	next *element
}

type Stack struct {
	sync.Mutex

	top  *element
	size int
}

func New() *Stack {
	s := new(Stack)

	return s
}

func (s *Stack) Push(data interface{}) {
	s.Lock()
	defer s.Unlock()

	ele := new(element)
	ele.data, ele.next = data, s.top
	s.top = ele

	s.size++
}

func (s *Stack) Pop() (data interface{}) {
	s.Lock()
	defer s.Unlock()

	if s.top == nil {
		return nil
	}

	data, s.top = s.top.data, s.top.next
	s.size--
	return
}

func (s *Stack) Size() int {
	s.Lock()
	defer s.Unlock()

	return s.size
}

func (s *Stack) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()

	return s.size == 0
}
