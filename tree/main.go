package main

import "fmt"

func main() {
	var t Tree

	t.Insert(5)
	t.Insert(1)
	t.Insert(6)
	t.Insert(3)
	t.Insert(8)

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
