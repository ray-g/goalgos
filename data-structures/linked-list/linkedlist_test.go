package linkedlist

import "testing"

func makeList(size int) *List {
	l := New()
	for i := 0; i < size; i++ {
		n := i
		l.PushBack(n)
	}
	return l
}

func TestListLength(t *testing.T) {
	if makeList(0).Length() != 0 {
		t.Error()
	}

	if makeList(3).Length() != 3 {
		t.Error()
	}
}

func TestListIsEmpty(t *testing.T) {
	if !makeList(0).IsEmpty() {
		t.Error()
	}

	if makeList(3).IsEmpty() {
		t.Error()
	}
}

func TestListFront(t *testing.T) {
	if makeList(0).Front() != nil {
		t.Error()
	}

	if makeList(3).Front().Value != 0 {
		t.Error()
	}
}

func TestListBack(t *testing.T) {
	if makeList(0).Back() != nil {
		t.Error()
	}

	if makeList(3).Back().Value != 2 {
		t.Error()
	}
}

func TestListPushFront(t *testing.T) {
	l := makeList(0)
	l.PushFront(4)
	if l.Front().Value != 4 {
		t.Error()
	}

	l.PushFront(5)
	if l.Front().Value != 5 {
		t.Error()
	}
}

func TestListPushBack(t *testing.T) {
	l := makeList(0)
	l.PushBack(3)
	if l.Back().Value != 3 {
		t.Error()
	}

	l.PushBack(6)
	if l.Back().Value != 6 {
		t.Error()
	}
}

func TestListInsertAt(t *testing.T) {
	l := makeList(0)
	if l.InsertAt(1, 1) ||
		l.InsertAt(2, -1) {
		t.Error()
	}

	if !l.InsertAt(4, 0) &&
		l.Front().Value != 4 {
		t.Error()
	}

	if !l.InsertAt(5, 0) &&
		l.Front().Value != 5 {
		t.Error()
	}

	if !l.InsertAt(7, 1) &&
		l.Back().Value != 7 {
		t.Error()
	}

	if !l.InsertAt(9, 1) &&
		l.Front().Next.Value != 9 {
		t.Error()
	}

	if !l.InsertAt(4, 2) &&
		l.Front().Next.Next.Value != 4 {
		t.Error()
	}
}

func TestListRemoveAt(t *testing.T) {
	l := makeList(0)

	v, ok := l.RemoveAt(-1)
	if ok ||
		v != nil {
		t.Error()
	}

	v, ok = l.RemoveAt(0)
	if ok ||
		v != nil {
		t.Error()
	}

	l.PushBack(4)

	v, ok = l.RemoveAt(2)
	if ok ||
		v != nil ||
		l.Length() != 1 {
		t.Error()
	}

	v, ok = l.RemoveAt(0)
	if !ok ||
		v != 4 ||
		l.Length() != 0 {
		t.Error()
	}

	l = makeList(7)

	v, ok = l.RemoveAt(0)
	if !ok ||
		v != 0 ||
		l.Length() != 6 {
		t.Error()
	}

	v, ok = l.RemoveAt(5)
	if !ok ||
		v != 6 ||
		l.Length() != 5 {
		t.Error()
	}

	v, ok = l.RemoveAt(1)
	if !ok ||
		v != 2 ||
		l.Length() != 4 {
		t.Error()
	}

	v, ok = l.RemoveAt(2)
	if !ok ||
		v != 4 ||
		l.Length() != 3 {
		t.Error()
	}
}

func TestListRemove(t *testing.T) {
	l := makeList(0)

	v, ok := l.Remove(0)
	if ok ||
		v != nil {
		t.Error()
	}

	l = makeList(7)
	v, ok = l.Remove(3)
	if !ok ||
		v != 3 ||
		l.Length() != 6 {
		t.Error()
	}

	v, ok = l.Remove(9)
	if ok ||
		v != nil ||
		l.Length() != 6 {
		t.Error()
	}
}

func TestListFind(t *testing.T) {
	l := makeList(0)

	if l.Find(0) != -1 {
		t.Error()
	}

	l = makeList(7)
	if l.Find(9) != -1 {
		t.Error()
	}

	if l.Find(3) != 3 {
		t.Error()
	}

	l.PushBack(3)
	if l.Find(3) != 3 {
		t.Error()
	}

	l.PushFront(3)
	if l.Find(3) != 0 {
		t.Error()
	}
}

func TestListGetAt(t *testing.T) {
	l := makeList(0)

	v, ok := l.GetAt(0)
	if ok ||
		v != nil {
		t.Error()
	}

	l = makeList(7)

	v, ok = l.GetAt(-1)
	if ok ||
		v != nil {
		t.Error()
	}

	v, ok = l.GetAt(8)
	if ok ||
		v != nil {
		t.Error()
	}

	v, ok = l.GetAt(3)
	if !ok ||
		v != 3 {
		t.Error()
	}
}

func TestListConcat(t *testing.T) {
	l1 := makeList(3)
	l2 := makeList(5)

	l1.Concat(l2)

	if l1.Length() != 8 ||
		l1.Front().Value != 0 ||
		l1.Back().Value != 4 {
		t.Error()
	}
}

func TestListClear(t *testing.T) {
	l := makeList(3)
	l.Clear()
	if l.Front() != nil ||
		l.Back() != nil ||
		l.Length() != 0 {
		t.Error()
	}
}

func TestListMap(t *testing.T) {
	l := makeList(0)
	count := 0

	l.Map(func(n *Node) {
		count++
	})

	if count != 0 {
		t.Error()
	}

	l = makeList(5)
	count = 0
	sum := 0
	var s [5]*Node
	ok := true

	l.Map(func(n *Node) {
		if n.Value != count {
			ok = false
		}
		s[count] = n
		count++
	})

	for i := 0; i < count; i++ {
		sum += s[i].Value.(int)
	}

	if !ok ||
		count != 5 ||
		sum != 10 {
		t.Error()
	}
}

func TestListEach(t *testing.T) {
	l := makeList(0)
	count := 0

	l.Each(func(n *Node) {
		count++
	})

	if count != 0 {
		t.Error()
	}

	l = makeList(5)
	count = 0
	sum := 0
	var s [5]*Node
	ok := true

	l.Each(func(n *Node) {
		if n.Value != count {
			ok = false
		}
		s[count] = n
		count++
	})

	for i := 0; i < count; i++ {
		sum += s[i].Value.(int)
	}

	if !ok ||
		count != 5 ||
		sum != 10 {
		t.Error()
	}
}
