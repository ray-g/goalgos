package binarytree

import (
	"goalgos/data-structures/queue"
	"goalgos/data-structures/stack"
)

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

func NewRoot(root *Node) *Tree {
	t := new(Tree)
	t.root = root
	return t
}

func (t *Tree) Add(value interface{}) {
	node := makeNode(value)
	if t.root == nil {
		t.root = node
	} else {
		added := false
		t.BreadthFirstSearch(func(n *Node) bool {
			if added {
				return false
			}
			if n.Left == nil {
				node.Parent = n
				n.Left = node
				added = true
			} else if n.Right == nil {
				node.Parent = n
				n.Right = node
				added = true
			}
			return !added
		})
	}
}

func (t *Tree) Delete(value interface{}) bool {
	deleted := false
	t.BreadthFirstSearch(func(n *Node) bool {
		if n.Value == value {
			if n.Parent != nil {
				if n.Parent.Left == n {
					n.Parent.Left = nil
				} else {
					n.Parent.Right = nil
				}
				n.Parent = nil
			} else { // Root node
				t.root = nil
			}
			deleted = true
		}
		return !deleted
	})
	return deleted
}

func (t *Tree) BreadthFirstSearch(f func(n *Node) bool) {
	if t.root == nil {
		return
	}

	searchQ := queue.New()
	searchQ.EnQueue(t.root)

	for !searchQ.IsEmpty() {
		node := searchQ.DeQueue().(*Node)
		if node.Left != nil {
			searchQ.EnQueue(node.Left)
		}

		if node.Right != nil {
			searchQ.EnQueue(node.Right)
		}

		if node == nil || !f(node) {
			break
		}
	}
}

func (t *Tree) DepthFirstSearch(f func(n *Node) bool) {
	if t.root == nil {
		return
	}

	searchS := stack.New()
	searchS.Push(t.root)

	for !searchS.IsEmpty() {
		node := searchS.Pop().(*Node)
		if node.Right != nil {
			searchS.Push(node.Right)
		}

		if node.Left != nil {
			searchS.Push(node.Left)
		}

		if node == nil || !f(node) {
			break
		}
	}
}
