package tree

import "fmt"

// Node of tree
type Node struct {
	Key   interface{}
	Value interface{}
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
	Less func(x, y interface{}) bool
}

// Add value to tree, return root node
func (tree *Tree) Add(key, value interface{}) *Node {
	tree.root = tree.add(key, value, tree.root)
	tree.n++
	return tree.root
}

func (tree *Tree) add(key, value interface{}, root *Node) *Node {
	if root == nil {
		return &Node{Key: key, Value: value}
	}
	if tree.Less(key, root.Key) {
		root.left = tree.add(key, value, root.left)
	} else {
		root.right = tree.add(key, value, root.right)
	}
	return root
}

// Search return Node for value
func (tree *Tree) Search(key interface{}) *Node {
	return tree.search(key, tree.root)
}
func (tree *Tree) search(key interface{}, root *Node) *Node {
	if root == nil {
		return nil
	} else if tree.Less(key, root.Key) {
		return tree.search(key, root.left)
	} else {
		return tree.search(key, root.right)
	}
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
func New(less func(x, y interface{}) bool) *Tree {
	return &Tree{
		Less: less,
	}
}
