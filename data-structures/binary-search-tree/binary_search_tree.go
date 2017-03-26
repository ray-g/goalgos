package bstree

// https://en.wikipedia.org/wiki/Binary_search_tree

type Node struct {
	Value       interface{}
	Parent      *Node
	Left, Right *Node
}

type BSTree struct {
	root    *Node
	compare func(*Node, *Node) int
}

func New(f func(*Node, *Node) int) *BSTree {
	t := new(BSTree)
	t.compare = f
	return t
}

func NewInt() *BSTree {
	t := new(BSTree)
	t.compare = func(a, b *Node) (r int) {
		r = 0
		if a.Value.(int) < b.Value.(int) {
			r = -1
		} else if a.Value.(int) > b.Value.(int) {
			r = 1
		}
		return
	}
	return t
}

func (t *BSTree) Insert(value interface{}) bool {
	return true
}

func (t *BSTree) Delete(value interface{}) bool {
	return true
}

func (t *BSTree) Find(value interface{}) *Node {
	return nil
}

func (t *BSTree) Balance() {
}

func (t *BSTree) IsBST() bool {
	return true
}

func (t *BSTree) Each(f func(n *Node) bool) {
}
