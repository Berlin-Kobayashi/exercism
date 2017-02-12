// Package tree implements a function for creating a tree structure from abstracted records.
package tree

import (
	"errors"
	"sort"
	"strconv"
)

const testVersion = 4

// Record represents an abstracted record in a tree.
type Record struct {
	ID, Parent int
}

type byID []Record

func (a byID) Len() int {
	return len(a)
}

func (a byID) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byID) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}

// Node represents a node in a tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a tree structure from abstracted records.
func Build(records []Record) (*Node, error) {
	sort.Sort(byID(records))

	if len(records) == 0 {
		return nil, nil
	}

	if !isValidRootRecord(records[0]) {
		return nil, errors.New("No valid root found")
	}

	root := &Node{ID: records[0].ID}
	nodes := []*Node{root}

	for i := 1; i < len(records); i++ {
		record := records[i]

		if record.ID != i {
			return nil, errors.New("Index is not properly incremented at index " + strconv.Itoa(i))
		}
		if !isValidRecord(record) {
			return nil, errors.New("Invalid record with index " + strconv.Itoa(i))
		}

		node := &Node{ID: record.ID}
		nodes[record.Parent].Children = append(nodes[record.Parent].Children, node)
		nodes = append(nodes, node)
	}

	return root, nil
}

func isValidRootRecord(root Record) bool {
	return root.ID == 0 && root.Parent == 0
}

func isValidRecord(record Record) bool {
	return record.ID > record.Parent
}
