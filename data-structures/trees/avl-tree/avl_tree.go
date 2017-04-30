package avltree

import (
	"fmt"
	"reflect"
	"strings"

	number "github.com/ray-g/goalgos/math/number-theory"
)

// https://en.wikipedia.org/wiki/AVL_tree

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

type AVLTree struct {
	root *Node
	typ  reflect.Type
	size int
}

func New(typ reflect.Type) *AVLTree {
	t := new(AVLTree)
	t.typ = typ
	return t
}

func (t *AVLTree) insertNode(node, parent *Node) bool {
	if node.less(parent) {
		// To Left
		if parent.Left == nil {
			node.Parent = parent
			parent.Left = node
			t.size++
			return true
		}
		if t.insertNode(node, parent.Left) {
			t.balance(parent)
			return balanceFactor(parent) != 0
		}
	}

	if parent.less(node) {
		// To Right
		if parent.Right == nil {
			node.Parent = parent
			parent.Right = node
			t.size++
			return true
		}
		if t.insertNode(node, parent.Right) {
			t.balance(parent)
			return balanceFactor(parent) != 0
		}
	}

	return false
}

func (t *AVLTree) Insert(value interface{}) bool {
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

	if t.insertNode(node, t.root) {
		return true
	}
	return false
}

func (t *AVLTree) Delete(value interface{}) bool {
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
			t.balance(parent)
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

			t.balance(parent)

			cut.dfsNext(func(n *Node) bool {
				t.size--
				t.Insert(n.Value)
				t.size++
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

func (t *AVLTree) Find(value interface{}) *Node {
	if t.typ != reflect.TypeOf(value) {
		return nil
	}

	return findNode(makeNode(value), t.root)
}

func (t *AVLTree) Size() int {
	return t.size
}

func (t *AVLTree) IsBST() bool {
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

func (t *AVLTree) Each(f func(n *Node) bool) {
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

func (t *AVLTree) Depth() int {
	return depth(t.root)
}

func balanceFactor(n *Node) int {
	return depth(n.Right) - depth(n.Left)
}

func (t *AVLTree) BalanceFactor() int {
	return balanceFactor(t.root)
}

func (t *AVLTree) balance(n *Node) {
	bal := balanceFactor(n)
	parent := n.Parent
	bRoot := parent == nil
	bLeft := false
	if !bRoot && parent.Left == n {
		bLeft = true
	}
	var pivot *Node
	if bal < -1 || bal > 1 {
		switch {
		case bal == -2 && balanceFactor(n.Left) == -1:
			pivot = rotateRight(n, n.Left)
		case bal == 2 && balanceFactor(n.Right) == 1:
			pivot = rotateLeft(n, n.Right)
		case bal == -2 && balanceFactor(n.Left) == 1:
			pivot = rotateLeftRight(n, n.Left.Right)
		case bal == 2 && balanceFactor(n.Right) == -1:
			pivot = rotateRightLeft(n, n.Right.Left)
		}

		if pivot != nil {
			pivot.Parent = parent
			if !bRoot {
				if bLeft {
					parent.Left = pivot
				} else {
					parent.Right = pivot
				}
			} else {
				t.root = pivot
			}
		} else {
			fmt.Println("No rotate?")
		}
	}
}

// Rotate functions
// root is the imbalance node
// pivot is the root node after re-balance

// Left-Left imbalance
//            r
//          /   \                   p
//         p     a                /   \
//       /  \          =>        n     r
//      n    b                  / \   / \
//     / \                     d   c b   a
//    d   c
func rotateRight(root, pivot *Node) *Node {
	child := pivot.Right

	pivot.Right = root
	root.Parent = pivot

	if child != nil {
		child.Parent = root
	}
	root.Left = child

	return pivot
}

// Right-Right imbalance
//      r
//    /   \                   p
//   a     p                /   \
//       /  \       =>     r     n
//      b    n            / \   / \
//          / \          a   b c   d
//         c   d
func rotateLeft(root, pivot *Node) *Node {
	child := pivot.Left

	pivot.Left = root
	root.Parent = pivot

	if child != nil {
		child.Parent = root
	}
	root.Right = child

	return pivot
}

// Left-Right imbalance
//            r                     r
//          /   \                 /   \                 p
//        n(r1)  a               p     a              /   \
//       /  \          =>       / \         =>       n     r
//      b    p(p1)             n   d                / \   / \
//          / \               / \                  b   c d   a
//         c   d             b   c
func rotateLeftRight(root, pivot *Node) *Node {
	rotateLeft(root.Left, pivot)
	pivot.Parent = root
	root.Left = pivot
	rotateRight(root, pivot)
	return pivot
}

// Right-Left imbalance
//     r                     r
//   /   \                 /   \                 p
//  a     n(r1)           a     p              /   \
//       /  \      =>          / \      =>    r     n
//     p(p1) b                d   n          / \   / \
//    / \                        / \        a   d c   b
//   d   c                      c   b
func rotateRightLeft(root, pivot *Node) *Node {
	rotateRight(root.Right, pivot)
	pivot.Parent = root
	root.Right = pivot
	rotateLeft(root, pivot)
	return pivot
}

// Dump dumps the structure of the subtree starting at node n
// including node search values and balance factors.
// Parameter i sets the line indent.
// lr is a prefix denoting the left or the right child, respectively.
func (n *Node) Dump(i int, lr string) {
	if n == nil {
		return
	}

	indent := ""
	if i > 0 {
		indent = strings.Repeat("|  ", (i-1)) + "+" + lr + ":"
	}
	fmt.Printf("%s%v[%d]\n", indent, n.Value, balanceFactor(n))
	n.Left.Dump(i+1, "L")
	n.Right.Dump(i+1, "R")
}

func (t *AVLTree) Dump() {
	t.root.Dump(0, "")
}

func (t *AVLTree) MinNode() *Node {
	node := t.root
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (t *AVLTree) MaxNode() *Node {
	node := t.root
	for node.Right != nil {
		node = node.Right
	}
	return node
}
