package stack

import "sync"

// StackS is slice implemented stack
type StackS struct {
	stack []interface{}
	lock  *sync.Mutex
	size  int
}

func NewST() *StackS {
	s := &StackS{}
	s.stack = make([]interface{}, 0)
	s.lock = &sync.Mutex{}

	return s
}

func (s *StackS) Push(ele interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	top := make([]interface{}, 1)
	top[0] = ele
	s.stack = append(top, s.stack...)

	s.size++
}

func (s *StackS) Pop() (ele interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		return nil
	}

	ele, s.stack = s.stack[0], s.stack[1:]
	s.size--
	return
}

func (s *StackS) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.size
}

func (s *StackS) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.size == 0
}
