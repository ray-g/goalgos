package btree

import "testing"

func makeBTree(size int) *Tree {
	t := New()
	for i := 0; i < size; i++ {
		t.Add(i)
	}
	return t
}

// Sample Tree
//                                       0
//                      _______________/   \_________________
//                     1                                      2
//            _______/   \_______                     ______/   \________
//           3                   4                   5                   6
//       __/   \__           __/   \_            __/   \_            __/   \_
//      7         8         9        10        11        12        13        14
//     / \      /   \     /   \     /   \     /   \     /   \     /   \     /   \
//   15   16   17   18   19   20   21   22   23   24   25   26   27   28   29   30
//   /
// 31
func sampleTree() *Tree {
	return makeBTree(32)
}

func TestTreeAdd(t *testing.T) {
	tree := sampleTree()
	num := []int{40, 41}

	verify := func() {
		for _, val := range num {
			tree.Add(val)

			found := false
			tree.DFSRecursive(func(n *Node) bool {
				if n.Value.(int) == val {
					found = true
					return true
				}
				return false
			})
			if !found {
				t.Error()
			}
		}
	}

	verify()

	tree = makeBTree(0)
	tree.Add(1)
	count := 0
	tree.BFS(func(n *Node) bool {
		count++
		return false
	})
	if count != 1 {
		t.Error()
	}
}

func TestTreeInsert(t *testing.T) {
	tree := makeBTree(0)
	if tree.Insert(0, 0, true) {
		t.Error()
	}

	tree = sampleTree()
	tree.Insert(40, 4, true)
	node := tree.Find(40)
	if node == nil ||
		node.Value.(int) != 40 ||
		node.Parent.Value.(int) != 4 ||
		node.Parent.Left != node ||
		node.Left.Value.(int) != 9 {
		t.Error()
	}
	tree.Insert(41, 16, true)
	node = tree.Find(41)
	if node == nil ||
		node.Value.(int) != 41 ||
		node.Parent.Value.(int) != 16 ||
		node.Parent.Left != node {
		t.Error()
	}
	tree.Insert(42, 16, false)
	node = tree.Find(42)
	if node == nil ||
		node.Value.(int) != 42 ||
		node.Parent.Value.(int) != 16 ||
		node.Parent.Right != node {
		t.Error()
	}
	tree.Insert(43, 4, false)
	node = tree.Find(43)
	if node == nil ||
		node.Value.(int) != 43 ||
		node.Parent.Value.(int) != 4 ||
		node.Parent.Right != node ||
		node.Right.Value.(int) != 10 {
		t.Error()
	}
}

func TestTreeDelete(t *testing.T) {
	tree := sampleTree()

	if !tree.Delete(32) {
		t.Error()
	}

	if tree.Delete(0) {
		t.Error()
	}

	if !tree.Delete(31) ||
		tree.Find(31) != nil ||
		tree.Find(15).Left != nil {
		t.Error()
	}

	if !tree.Delete(30) ||
		tree.Find(30) != nil ||
		tree.Find(14).Right != nil {
		t.Error()
	}

	tree.Delete(18)
	if !tree.Delete(8) ||
		tree.Find(8) != nil ||
		tree.Find(3).Right.Value.(int) != 17 ||
		tree.Find(17).Parent.Value.(int) != 3 {
		t.Error()
	}

	tree.Delete(19)
	if !tree.Delete(9) ||
		tree.Find(9) != nil ||
		tree.Find(4).Left.Value.(int) != 20 ||
		tree.Find(20).Parent.Value.(int) != 4 {
		t.Error()
	}

	tree.Delete(24)
	if !tree.Delete(11) ||
		tree.Find(11) != nil ||
		tree.Find(5).Left.Value.(int) != 23 ||
		tree.Find(23).Parent.Value.(int) != 5 {
		t.Error()
	}

	tree.Delete(25)
	if !tree.Delete(12) ||
		tree.Find(12) != nil ||
		tree.Find(5).Right.Value.(int) != 26 ||
		tree.Find(26).Parent.Value.(int) != 5 {
		t.Error()
	}

	tree = makeBTree(2)
	if !tree.Delete(0) ||
		tree.root.Value.(int) != 1 ||
		tree.root.Parent != nil {
		t.Error()
	}

	tree = makeBTree(1)
	tree.Insert(1, 0, false)
	if !tree.Delete(0) ||
		tree.root.Value.(int) != 1 ||
		tree.root.Parent != nil {
		t.Error()
	}
}

func TestTreeFind(t *testing.T) {
	tree := makeBTree(0)
	n := tree.Find(4)
	if n != nil {
		t.Error()
	}

	tree = sampleTree()
	n = tree.Find(32)
	if n != nil {
		t.Error()
	}

	n = tree.Find(17)
	if n == nil || n.Value.(int) != 17 {
		t.Error()
	}
}

func TestTreeBFS(t *testing.T) {
	tree := makeBTree(0)
	count := 0
	tree.BFS(func(n *Node) bool {
		count++
		return false
	})
	if count != 0 {
		t.Error()
	}

	tree = sampleTree()
	index := 0
	good := true
	tree.BFS(func(n *Node) bool {
		if n.Value.(int) != index {
			good = false
		}
		index++
		return false
	})
	if index != 32 || !good {
		t.Error()
	}
}

func BenchmarkBFS(b *testing.B) {
	tree := makeBTree(1000)

	for i := 0; i < b.N; i++ {
		tree.BFS(func(n *Node) bool {
			return false
		})
	}
}

func TestTreeDFS(t *testing.T) {
	tree := makeBTree(0)
	count := 0
	tree.DFSStack(func(n *Node) bool {
		count++
		return false
	})
	if count != 0 {
		t.Error()
	}

	tree.DFSRecursive(func(n *Node) bool {
		count++
		return false
	})
	if count != 0 {
		t.Error()
	}

	tree = sampleTree()

	found := false
	tree.DFSStack(func(n *Node) bool {
		if n.Value.(int) == 15 {
			found = true
		}
		return found
	})
	if !found {
		t.Error()
	}

	found = false
	tree.DFSRecursive(func(n *Node) bool {
		if n.Value.(int) == 15 {
			found = true
		}
		return found
	})
	if !found {
		t.Error()
	}

	index := 0
	good := true
	order := []int{0, 1, 3, 7, 15, 31, 16, 8, 17, 18, 4, 9, 19, 20, 10, 21, 22, 2, 5, 11, 23, 24, 12, 25, 26, 6, 13, 27, 28, 14, 29, 30}
	tree.DFSStack(func(n *Node) bool {
		if n.Value.(int) != order[index] {
			good = false
		}
		index++
		return false
	})
	if index != 32 || !good {
		t.Error()
	}

	index = 0
	good = true
	tree.DFSRecursive(func(n *Node) bool {
		if n.Value.(int) != order[index] {
			good = false
		}
		index++
		return false
	})
	if index != 32 || !good {
		t.Error()
	}
}

func BenchmarkDFSStack(b *testing.B) {
	tree := makeBTree(1000)

	for i := 0; i < b.N; i++ {
		tree.DFSStack(func(n *Node) bool {
			return false
		})
	}
}

func BenchmarkDFSRecursive(b *testing.B) {
	tree := makeBTree(1000)

	for i := 0; i < b.N; i++ {
		tree.DFSRecursive(func(n *Node) bool {
			return false
		})
	}
}
