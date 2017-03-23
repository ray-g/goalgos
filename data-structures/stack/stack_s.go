package stack

import "sync"

// StackS is slice implemented stack
type StackS struct {
	sync.Mutex

	stack []interface{}
	size  int
}

func NewST() *StackS {
	s := &StackS{}
	s.stack = make([]interface{}, 0)

	return s
}

func (s *StackS) Push(ele interface{}) {
	s.Lock()
	defer s.Unlock()

	top := make([]interface{}, 1)
	top[0] = ele
	s.stack = append(top, s.stack...)

	s.size++
}

func (s *StackS) Pop() (ele interface{}) {
	s.Lock()
	defer s.Unlock()

	if s.size == 0 {
		return nil
	}

	ele, s.stack = s.stack[0], s.stack[1:]
	s.size--
	return
}

func (s *StackS) Size() int {
	s.Lock()
	defer s.Unlock()

	return s.size
}

func (s *StackS) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()

	return s.size == 0
}
