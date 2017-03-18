package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()

	if !s.IsEmpty() ||
		s.size != 0 ||
		s.Size() != 0 {
		t.Error()
	}

	s.Push(0)
	s.Push(1)
	s.Push(2)

	if s.Size() != 3 {
		t.Error()
	}

	e2 := s.Pop()

	if e2 != 2 || s.Size() != 2 {
		t.Error()
	}

	e1 := s.Pop()

	if e1 != 1 {
		t.Error()
	}

	e0 := s.Pop()

	if e0 != 0 || s.Size() != 0 {
		t.Error()
	}

	ee := s.Pop()

	if ee != nil || s.Size() != 0 {
		t.Error()
	}
}
