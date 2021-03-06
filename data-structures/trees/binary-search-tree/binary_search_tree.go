package bstree

import (
	"reflect"

	number "github.com/ray-g/goalgos/math/number-theory"
)

// https://en.wikipedia.org/wiki/Binary_search_tree

type Item interface {
	Less(other interface{}) bool
}

type Node struct {
	Value       Item
	Parent      *Node
	Left, Right *Node
}

func (n *Node) less(than *Node) bool {
	return n.Value.Less(than.Value)
}

func makeNode(value interface{}) *Node {
	if _, ok := value.(Item); ok {
		return &Node{Value: value.(Item)}
	}
	return nil
}

type BSTree struct {
	root *Node
	typ  reflect.Type
	size int
}

func New(typ reflect.Type) *BSTree {
	t := new(BSTree)
	t.typ = typ
	return t
}

func insertNode(node, parent *Node) bool {
	if node.less(parent) {
		// To Left
		if parent.Left == nil {
			node.Parent = parent
			parent.Left = node
			return true
		}
		return insertNode(node, parent.Left)
	}

	if parent.less(node) {
		// To Right
		if parent.Right == nil {
			node.Parent = parent
			parent.Right = node
			return true
		}
		return insertNode(node, parent.Right)
	}

	return false
}

func (t *BSTree) Insert(value interface{}) bool {
	if t.typ != reflect.TypeOf(value) {
		return false
	}

	node := makeNode(value)
	if node == nil {
		return false
	}

	if t.root == nil {
		t.root = node
		t.size++
		return true
	}

	if insertNode(node, t.root) {
		t.size++
		return true
	}
	return false
}

func (t *BSTree) Delete(value interface{}) bool {
	if t.typ != reflect.TypeOf(value) {
		return false
	}

	node := findNode(makeNode(value), t.root)
	if node != nil {
		if node.Left == nil && node.Right == nil {
			if node.Parent == nil {
				t.root = nil
			} else {
				if node.Parent.Left == node {
					node.Parent.Left = nil
				} else {
					node.Parent.Right = nil
				}
				node.Parent = nil
			}
		}

		if node.Left != nil && node.Right == nil ||
			node.Left == nil && node.Right != nil {
			var child *Node
			if node.Left != nil {
				child = node.Left
			} else {
				child = node.Right
			}

			parent := node.Parent
			child.Parent = parent

			if parent.Left == node {
				parent.Left = child
			} else {
				parent.Right = child
			}
			node.Parent = nil
		}

		if node.Left != nil && node.Right != nil {
			var child, cut *Node
			parent := node.Parent
			if depth(node.Left) < depth(node.Right) {
				child = node.Left
				cut = node.Right
				node.Left = nil
			} else {
				child = node.Right
				cut = node.Left
				node.Right = nil
			}
			child.Parent = parent
			if parent.Left == node {
				parent.Left = child
			} else {
				parent.Right = child
			}

			node.Parent = nil

			cut.dfsNext(func(n *Node) bool {
				t.Insert(n.Value)
				return false
			})
		}
		t.size--
	}

	return true
}

func findNode(val, node *Node) *Node {
	if val == nil || node == nil {
		return nil
	}
	if val.less(node) && node.Left != nil {
		return findNode(val, node.Left)
	}

	if node.less(val) && node.Right != nil {
		return findNode(val, node.Right)
	}

	if !val.less(node) && !node.less(val) {
		return node
	}

	return nil
}

func (t *BSTree) Find(value interface{}) *Node {
	if t.typ != reflect.TypeOf(value) {
		return nil
	}

	return findNode(makeNode(value), t.root)
}

func (t *BSTree) Size() int {
	return t.size
}

func (t *BSTree) IsBST() bool {
	isBst := true
	t.Each(func(n *Node) bool {
		if n.Left != nil && n.less(n.Left) ||
			n.Right != nil && n.Right.less(n) {
			isBst = false
		}
		return false
	})
	return isBst
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

func (t *BSTree) Each(f func(n *Node) bool) {
	if t.root == nil {
		return
	}

	t.root.dfsNext(f)
}

func depth(n *Node) int {
	if n == nil {
		return 0
	}

	return 1 + number.Max(depth(n.Left), depth(n.Right))
}

func (t *BSTree) Depth() int {
	return depth(t.root)
}

func balanceFactor(n *Node) int {
	return depth(n.Right) - depth(n.Left)
}

func (t *BSTree) BalanceFactor() int {
	return balanceFactor(t.root)
}
