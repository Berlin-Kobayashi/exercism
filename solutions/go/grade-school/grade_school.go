// Package school implements an API for managing students of a school for different grades.
package school

import (
	"sort"
)

// Grade represents a grade with its students.
type Grade struct {
	grade    int
	students []string
}

// School represents a school with its grades.
type School map[int]Grade

type grades []Grade

// New creates a new school and returns a pointer to it.
func New() *School {
	school := School(make(map[int]Grade))

	return &school
}

// Add adds the given student to the given grade.
func (s *School) Add(student string, grade int) {
	if _, gradeExists := (*s)[grade]; !gradeExists {
		(*s)[grade] = Grade{grade, []string{}}
	}

	students := append(s.Grade(grade), student)
	(*s)[grade] = Grade{grade, students}
}

// Grade returns all students in the given grade in ascending order.
func (s *School) Grade(grade int) []string {
	sort.Sort(sort.StringSlice((*s)[grade].students))

	return (*s)[grade].students
}

// Enrollment returns all grades of the school in ascending order.
func (s *School) Enrollment() []Grade {
	gradeSlice := make([]Grade, len(map[int]Grade(*s)))

	i := 0
	for _, grade := range *s {
		gradeSlice[i] = grade
		sort.Sort(sort.StringSlice(gradeSlice[i].students))
		i++
	}

	sort.Sort(grades(gradeSlice))

	return gradeSlice
}

func (g grades) Len() int {
	return len(g)
}

func (g grades) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g grades) Less(i, j int) bool {
	return g[i].grade < g[j].grade
}
