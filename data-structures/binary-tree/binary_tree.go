package btree

// https://en.wikipedia.org/wiki/Binary_tree

import (
	"github.com/ray-g/goalgos/data-structures/queue"
	"github.com/ray-g/goalgos/data-structures/stack"
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

// Add to the first nil leaf. First is according BFS
func (t *Tree) Add(value interface{}) {
	node := makeNode(value)
	if t.root == nil {
		t.root = node
	} else {
		added := false
		t.BFS(func(n *Node) bool {
			if n.Left == nil {
				node.Parent = n
				n.Left = node
				added = true
			} else if n.Right == nil {
				node.Parent = n
				n.Right = node
				added = true
			}
			return added
		})
	}
}

// Insert follow https://en.wikipedia.org/wiki/Binary_tree#Insertion
func (t *Tree) Insert(value interface{}, at interface{}, toLeft bool) bool {
	parent := t.Find(at)

	if parent != nil {
		node := makeNode(value)
		node.Parent = parent
		if toLeft {
			if parent.Left == nil {
				parent.Left = node
			} else {
				node.Left = parent.Left
				node.Left.Parent, parent.Left = node, node
			}
		} else {
			if parent.Right == nil {
				parent.Right = node
			} else {
				node.Right = parent.Right
				node.Right.Parent, parent.Right = node, node
			}
		}
		return true
	}
	return false
}

// Delete follow https://en.wikipedia.org/wiki/Binary_tree#Deletion
func (t *Tree) Delete(value interface{}) (deleted bool) {
	deleted = true
	node := t.Find(value)
	if node == nil {
		return
	} else if node.Left != nil && node.Right != nil {
		// Have both leaves, cannot delete
		deleted = false
	} else if node.Parent == nil {
		// Root
		if node.Left != nil {
			t.root = node.Left
		} else {
			t.root = node.Right
		}
		t.root.Parent = nil
	} else {
		parent := node.Parent
		if node.Left == nil && node.Right == nil {
			// Leaf, just delete it
			if parent.Left == node {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			node.Parent = nil
		} else if node.Left != nil {
			// Left is not nil
			node.Left.Parent = parent
			if parent.Left == node {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
			node.Parent = nil
		} else {
			// Right is not nil
			node.Right.Parent = parent
			if parent.Left == node {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
			node.Parent = nil
		}
	}
	return
}

func (t *Tree) Find(value interface{}) *Node {
	var node *Node
	t.BFS(func(n *Node) bool {
		if n.Value == value {
			node = n
			return true
		}
		return false
	})
	return node
}

// BFS Breadth First Search
func (t *Tree) BFS(f func(n *Node) bool) {
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

		if node == nil || f(node) {
			break
		}
	}
}

// DFS Depth First Search
func (t *Tree) DFSStack(f func(n *Node) bool) {
	if t.root == nil {
		return
	}

	t.dfsStack(f)
}

func (t *Tree) DFSRecursive(f func(n *Node) bool) {
	if t.root == nil {
		return
	}

	t.root.dfsNext(f)
}

func (t *Tree) dfsStack(f func(n *Node) bool) {
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

		if node == nil || f(node) {
			break
		}
	}
}

func (n *Node) dfsNext(f func(n *Node) bool) bool {
	if f(n) {
		return true
	}

	found := false
	if n.Left != nil {
		found = n.Left.dfsNext(f)
	}
	if !found && n.Right != nil {
		found = n.Right.dfsNext(f)
	}
	return found
}
