//Package binarysearchtree implements an API for working with binary search trees.
package binarysearchtree

// SearchTreeData represents a node in a binary search tree.
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst creates a new binary tree holding the given value.
func Bst(value int) *SearchTreeData {
	tree := SearchTreeData{
		left:  nil,
		data:  value,
		right: nil,
	}

	return &tree
}

// Insert inserts the given value into the binary search tree.
func (tree *SearchTreeData) Insert(value int) {
	if value <= tree.data {
		if tree.left == nil {
			tree.left = Bst(value)
		} else {
			tree.left.Insert(value)
		}
	} else {
		if tree.right == nil {
			tree.right = Bst(value)
		} else {
			tree.right.Insert(value)
		}
	}
}

// MapString executes the given function on each element in an ascending order and returns an string array of the results.
func (tree *SearchTreeData) MapString(function func(int) string) (result []string) {
	tree.walk(func(value int) {
		result = append(result, function(value))
	})

	return result
}

// MapInt executes the given function on each element in an ascending order and returns an integer array of the results.
func (tree *SearchTreeData) MapInt(function func(int) int) (result []int) {
	tree.walk(func(value int) {
		result = append(result, function(value))
	})

	return result
}

func (tree *SearchTreeData) walk(function func(int)) {
	if tree.left != nil || tree.right != nil {
		if tree.left != nil {
			tree.left.walk(function)
			function(tree.data)
		}
		if tree.right != nil {
			if tree.left == nil {
				function(tree.data)
			}
			tree.right.walk(function)
		}
	} else {
		function(tree.data)
	}
}
