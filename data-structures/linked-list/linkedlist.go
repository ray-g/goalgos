package linkedlist

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

func makeNode(value interface{}) *Node {
	return &Node{Value: value}
}

type List struct {
	front *Node
	back  *Node
	len   int
}

func New() *List {
	l := new(List)
	l.front = nil
	l.back = nil
	l.len = 0

	return l
}

func (l *List) Length() int {
	return l.len
}

func (l *List) IsEmpty() bool {
	return l.len == 0
}

func (l *List) Front() *Node {
	return l.front
}

func (l *List) Back() *Node {
	return l.back
}

func (l *List) PushFront(val interface{}) {
	node := makeNode(val)

	if l.len == 0 {
		l.front, l.back = node, node
	} else {
		l.front.Prev, node.Next = node, l.front
		l.front = node
	}

	l.len++
}

func (l *List) PushBack(val interface{}) {
	node := makeNode(val)

	if l.len == 0 {
		l.front, l.back = node, node
	} else {
		l.back.Next, node.Prev = node, l.back
		l.back = node
	}

	l.len++
}

func (l *List) InsertAt(val interface{}, index int) bool {
	if index < 0 || index > l.len {
		return false
	}

	if index == 0 || l.len == 0 {
		l.PushFront(val)
		return true
	}

	if index == l.len-1 {
		l.PushBack(val)
		return true
	}

	node := makeNode(val)
	at := l.front.Next
	for i := 1; i < index; i++ {
		at = at.Next
	}
	node.Next, node.Prev = at, at.Prev
	node.Prev.Next, node.Next.Prev = node, node
	l.len++
	return true
}

func (l *List) RemoveAt(index int) (val interface{}, _ bool) {
	if index < 0 || l.len == 0 || index > l.len {
		return nil, false
	}

	if l.len == 1 {
		val = l.front.Value
		l.front, l.back = nil, nil
	} else if index == 0 {
		val = l.front.Value
		l.front = l.front.Next
		l.front.Prev = nil
	} else if index == l.len-1 {
		val = l.back.Value
		l.back = l.back.Prev
		l.back.Next = nil
	} else {
		node := l.front.Next
		for i := 1; i < index; i++ {
			node = node.Next
		}
		node.Prev.Next, node.Next.Prev = node.Next, node.Prev
		val = node.Value
	}
	l.len--
	return val, true
}

func (l *List) Remove(val interface{}) (interface{}, bool) {
	index := l.Find(val)

	if index != -1 {
		return l.RemoveAt(index)
	}
	return nil, false
}

func (l *List) Find(val interface{}) int {
	if l.len == 0 {
		return -1
	}

	index, i := -1, 0
	l.Each(func(n *Node) {
		if n.Value == val && index == -1 {
			index = i
		}
		i++
	})
	return index
}

func (l *List) GetAt(index int) (interface{}, bool) {
	if index < 0 || l.len == 0 || index > l.len {
		return nil, false
	}

	node := l.front
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node.Value, true
}

func (l *List) Concat(ol *List) {
	l.back.Next, ol.front.Prev = ol.front, l.back
	l.back = ol.back
	l.len += ol.len
}

func (l *List) Clear() {
	l.len = 0
	l.front, l.back = nil, nil
}

func (l *List) Map(f func(n *Node)) {
	for node := l.front; node != nil; node = node.Next {
		n := makeNode(node.Value)
		f(n)
	}
}

func (l *List) Each(f func(n *Node)) {
	for node := l.front; node != nil; node = node.Next {
		f(node)
	}
}
