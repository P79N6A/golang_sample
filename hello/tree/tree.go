package tree

import "fmt"

// Node of tree
type Node struct {
	Value int
	left  *Node
	right *Node
}

// String of Node
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

// Tree is a binary search tree
type Tree struct {
	root *Node
	n    int
}

// Add value to tree, return root node
func (tree *Tree) Add(value int) *Node {
	tree.root = add(value, tree.root)
	tree.n++
	return tree.root
}

func add(value int, root *Node) *Node {
	node := Node{Value: value}
	if root == nil {
		return &node
	}
	if value < root.Value {
		root.left = add(value, root.left)
	} else {
		root.right = add(value, root.right)
	}
	return root
}

// Traverse a tree
func (tree *Tree) Traverse() {
	traverse(tree.root)
}

func traverse(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root)
	traverse(root.left)
	traverse(root.right)
}

// TraverseByLevel traverse tree by level
func (tree *Tree) TraverseByLevel() {
	if tree.root == nil {
		return
	}
	queue := []*Node{}
	queue = append(queue, tree.root)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		fmt.Println(n)
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
}

// New constuct a tree
func New(value ...int) *Tree {
	tree := Tree{}
	for _, v := range value {
		tree.Add(v)
	}
	return &tree
}
