package binarytree

type Node struct {
	Value       interface{}
	Parent      *Node
	Left, Right *Node
}

func makeNode(value interface{}) *Node {
	return &Node{Value: value}
}

type Tree struct {
	root *Node
}

func New() *Tree {
	t := new(Tree)
	return t
}

func (t *Tree) Add(value interface{}) {

}

func (t *Tree) Delete(value interface{}) bool {
	return false
}

func (t *Tree) BreadthFirstSearch(f func(n *Node)) {

}

func (t *Tree) DepthFirstSearch(f func(n *Node)) {

}
