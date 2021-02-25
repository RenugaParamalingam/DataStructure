package main

import "fmt"

func main() {
	var t Tree

	t.Insert(4)
	t.Insert(1)
	t.Insert(6)
	t.Insert(3)
	t.Insert(5)
	t.Insert(8)
	t.Insert(7)
	t.Insert(9)

	Delete(t.Root, 6)

	fmt.Println("Pre order traversal")
	t.PreOrder(t.Root)

	fmt.Println("\n Post order traversal")
	t.PostOrder(t.Root)

	fmt.Println("\n Inorder traversal")
	t.InOrder(t.Root)

	searchKey := 9

	fmt.Println("\n Is key ", searchKey, "present in tree: ", t.Search(searchKey))

	fmt.Println("Level order")
	t.PrintLevelOrder()

	fmt.Println("\n Get level order: ", t.GetLevelOrder())

	fmt.Println("max depth: ", MaxDepth(t.Root))
	fmt.Println("min depth: ", MinDepth(t.Root))

	invalidTree := Tree{}
	invalidTree.Root = &Node{Data: 1, Left: &Node{Data: 2}, Right: &Node{Data: 3}}

	fmt.Println("minimum value: ", t.MinValue())
	fmt.Println("maximum value: ", t.MaxValue())

	// t.Invert()
	mirror(t.Root)
	fmt.Println("\n Inorder traversal after mirror")
	t.InOrder(t.Root)
}

// Node holds structure for a node in binary tree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Tree holds structure of a tree
type Tree struct {
	Root *Node
}

// Insert inserts the node into binary search tree
func (t *Tree) Insert(data int) {
	if t.Root == nil {
		t.Root = &Node{Data: data}
		return
	}

	t.Root.insertNode(data)
}

func (n *Node) insertNode(data int) {
	if data <= n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: data}
		} else {
			n.Left.insertNode(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Data: data}
		} else {
			n.Right.insertNode(data)
		}
	}
}

// Delete removes the node from binary search tree.
func Delete(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.Data {
		node.Left = Delete(node.Left, key)
		return node
	}

	if key > node.Data {
		node.Right = Delete(node.Right, key)
		return node
	}

	if node.Left == nil && node.Right == nil {
		node = nil
		return nil
	}

	if node.Left == nil {
		node = node.Right
		return node
	}

	if node.Right == nil {
		node = node.Left
		return node
	}

	// left most in the right subtree
	leftMostRight := node.Right

	for {
		if leftMostRight != nil && leftMostRight.Left != nil {
			leftMostRight = leftMostRight.Left
		} else {
			break
		}
	}

	node.Data = leftMostRight.Data
	node.Right = Delete(node.Right, node.Data)

	return node
}

// PreOrder implements pre-order traversal in binary search tree
func (t *Tree) PreOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d -> ", node.Data)

		t.PreOrder(node.Left)
		t.PreOrder(node.Right)
	}
}

// PostOrder implements post-order traversal in binary search tree
func (t *Tree) PostOrder(node *Node) {
	if node != nil {
		t.PreOrder(node.Left)
		t.PreOrder(node.Right)

		fmt.Printf("%d -> ", node.Data)
	}
}

// InOrder implements in-order traversal in binary search tree
func (t *Tree) InOrder(node *Node) {
	if node != nil {
		t.PreOrder(node.Left)
		fmt.Printf("%d -> ", node.Data)
		t.PreOrder(node.Right)
	}
}

// Search returns true if search key is found in the binary search tree
func (t *Tree) Search(key int) bool {
	current := t.Root

	for current != nil {
		if key == current.Data {
			return true
		}

		if key < current.Data {
			current = current.Left
		}

		if key > current.Data {
			current = current.Right
		}
	}

	return false
}

// PrintLevelOrder prints tree at the level
func (t *Tree) PrintLevelOrder() {
	q := make([]*Node, 0)
	q = append(q, t.Root)

	for len(q) > 0 {
		temp := q[0]

		fmt.Print(temp.Data)

		if temp.Left != nil {
			q = append(q, temp.Left)
		}

		if temp.Right != nil {
			q = append(q, temp.Right)
		}

		q = q[1:]
	}
}

// GetLevelOrder returns the level order as a 2D array
func (t *Tree) GetLevelOrder() [][]int {
	var levels [][]int
	q := []*Node{t.Root}

	for len(q) > 0 {
		level := []int{}
		l := len(q)

		for i := 0; i < l; i++ {
			if q[0] != nil {
				level = append(level, q[0].Data)
				q = append(q, q[0].Left)
				q = append(q, q[0].Right)
			}
			q = q[1:]
		}

		levels = append(levels, level)
	}

	return levels[:len(levels)-1]
}

// MaxDepth returns the maximum depth of the tree.
func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinDepth returns minimum depth of the tree.
func MinDepth(root *Node) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}

	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}

	return min(MinDepth(root.Left), MinDepth(root.Right)) + 1
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//MinValue returns minimum value in the tree.
func (t *Tree) MinValue() int {
	if t.Root == nil {
		return -1
	}

	node := t.Root
	for node != nil {
		if node.Left == nil {
			return node.Data
		}
		node = node.Left
	}

	return -1
}

//MaxValue returns minimum value in the tree.
func (t *Tree) MaxValue() int {
	if t.Root == nil {
		return -1
	}

	node := t.Root
	for node != nil {
		if node.Right == nil {
			return node.Data
		}
		node = node.Right
	}

	return -1
}

// Invert inverts left and right of the binary tree.
func (t *Tree) Invert() {
	q := []*Node{t.Root}

	for len(q) > 0 {
		current := q[0]

		if current.Left != nil && current.Right != nil {
			temp := current.Left.Data
			current.Left.Data = current.Right.Data
			current.Right.Data = temp
		} else if current.Left == nil {
			current.Left = current.Right
			current.Right = nil
		} else if current.Right == nil {
			current.Right = current.Left
			current.Left = nil
		}

		if current.Left != nil {
			q = append(q, current.Left)
		}

		if current.Right != nil {
			q = append(q, current.Right)
		}

		q = q[1:]
	}
}

func mirror(node *Node) {
	if node == nil {
		return
	}

	mirror(node.Left)
	mirror(node.Right)

	temp := node.Left
	node.Left = node.Right
	node.Right = temp
}
