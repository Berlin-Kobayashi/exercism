// Package kindergarten implements a simple API for determining which plant belongs to which kindergarten student.
package kindergarten

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

const testVersion = 1

type Garden map[string][]string

// NewGarden creates a new garden based on the given diagram and the children names.
// Returns an error if there is a duplicate child name or if the diagram is invalid.
func NewGarden(diagram string, children []string) (*Garden, error) {
	sortedChildren := sortChildrenByName(children)

	err := validateSortedChildren(sortedChildren)
	if err != nil {
		return nil, err
	}

	diagramRows, err := getDiagramRows(diagram, len(sortedChildren))
	if err != nil {
		return nil, err
	}

	garden := make(Garden, len(children))
	columnCounter := 0
	for _, child := range sortedChildren {
		plants := make([]string, 4)
		plantCounter := 0
		for rowCounter := 0; rowCounter < len(diagramRows); rowCounter++ {
			for _, plantID := range diagramRows[rowCounter][columnCounter : columnCounter+2] {
				plants[plantCounter], err = getPlantByID(plantID)
				if err != nil {
					return nil, err
				}

				plantCounter++
			}
		}
		garden[child] = plants

		columnCounter += 2
	}

	return &garden, nil
}

func sortChildrenByName(children []string) []string {
	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	return sortedChildren
}

func validateSortedChildren(sortedChildren []string) error {
	previousChild := ""
	for _, child := range sortedChildren {
		if child == previousChild {
			return fmt.Errorf("Invalid sortedChildren given. Childrens must be unique but %s is not.", child)
		}
		previousChild = child
	}

	return nil
}

func getDiagramRows(diagram string, childrenAmount int) ([]string, error) {
	diagramRows := strings.Split(diagram, "\n")
	if len(diagramRows) != 3 {
		return nil, errors.New("Invalid diagram given. Diagram must have 3 lines.")
	}

	diagramRows = diagramRows[1:]

	for _, row := range diagramRows {
		if len(row) != childrenAmount*2 {
			return nil, errors.New("Invalid diagram given. Rows must have twice as much plants as there are children.")
		}
	}

	return diagramRows, nil
}

func getPlantByID(id rune) (string, error) {
	switch id {
	case 'G':
		return "grass", nil
	case 'C':
		return "clover", nil
	case 'R':
		return "radishes", nil
	case 'V':
		return "violets", nil
	default:
		return "", fmt.Errorf("Invalid plantID. Given: %c", id)

	}
}

// Plants returns the name of the plants of the child. Returns false if the child is unknown .
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]

	return plants, ok
}
