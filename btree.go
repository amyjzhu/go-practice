package main

import "fmt"

type BTree struct {
	value int
	left *BTree
	right *BTree
}

func (tree *BTree) print() {
	fmt.Println(tree.value)
	// ugh, I'm lazy
	tree.left.print()
	tree.right.print()
}

func insert_array(tree []int, target int) {
	// heaps

}

// might as well treat as a BST
// no duplicates allowed
func insert(tree *BTree, target int) {
	node := BTree{target, nil, nil}

	if tree == nil {
		tree = &node
		return
	}

	if (tree.value > target) {
		if tree.left != nil {
			tree.left = &node
		} else {
			insert(tree.left, target)
		}
	} else if (tree.value < target) {
		if tree.right != nil {
			tree.right = &node
		} else {
			insert(tree.right, target)
		}
	}
}

