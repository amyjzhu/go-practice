package main

import "fmt"

type BTree struct {
	value int
	left *BTree
	right *BTree
}

func (tree *BTree) print() {
	if (tree == nil) {
		return
	}
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
		if tree.left == nil {
			tree.left = &node
		} else {
			insert(tree.left, target)
		}
	} else if (tree.value < target) {
		if tree.right == nil {
			tree.right = &node
		} else {
			insert(tree.right, target)
		}
	}
}

func query(tree *BTree, target int) bool {

	if (tree == nil) {
		return false;
	}

	if (tree.value == target) {
		return true;
	} else if (tree.value > target) {
		return query(tree.left, target)
	} else if (tree.value < target) {
		return query(tree.right, target)
	}

	return false;
}

func deleteTree(tree *BTree, target int) {
	_deleteTree(tree, target, nil, false)
}
// can I just assign the new tree
// this is so ugly
func _deleteTree(tree *BTree, target int, parent *BTree, right bool) {
	// first find element
	if (target > tree.value) {
		_deleteTree(tree.right, target, tree, true)
	}

	if (target < tree.value) {
		_deleteTree(tree.left, target, tree, false)
	}

	if (target == tree.value) {
		// is it root or leaf?
		if (tree.left == nil && tree.right == nil) {
			if (right) {
				parent.right = nil
			} else {
				parent.left = nil // cannot use "nil" in BTree assignment...
				// left or right?
			}
		} else if (tree.left == nil) {
			// !!! need to use pointers here to assign correctly, looks like
			// I guess assigning the pointer to nil doesn't do anything since it's an argument
			*tree = *tree.right // possibly nil
		} else if (tree.right == nil) {
			*tree = *tree.left
		} else {
			reassign := find_leftmost(tree.right)
			tree.value = reassign
			_deleteTree(tree.right, reassign, tree, true)
		}
	}
}

func find_leftmost(tree *BTree) int {
	if tree == nil {
		return -1
	}

	var leftmost_so_far int;
	var leftTree = tree.left

	for leftTree != nil {
		leftmost_so_far = leftTree.value
		leftTree = leftTree.left
	}

	return leftmost_so_far
}