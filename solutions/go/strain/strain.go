// Package strain implements a set of filterable collections.
package strain

// Ints represents a collection of ints.
type Ints []int

// Keep returns a collection containing all elements for which the given callback returns true.
func (ints Ints) Keep(callback func(int) bool) Ints {
	if ints == nil {
		return nil
	}

	result := Ints{}
	for _, number := range ints {
		if callback(number) {
			result = append(result, number)
		}
	}

	return result
}

// Discard returns a collection containing all elements for which the given callback returns false.
func (ints Ints) Discard(callback func(int) bool) Ints {
	negatedCallback := func(v int) bool { return !callback(v) }

	return ints.Keep(negatedCallback)
}

// Lists represents a collection of int slices.
type Lists [][]int

// Keep returns a collection containing all elements for which the given callback returns true.
func (lists Lists) Keep(callback func([]int) bool) Lists {
	if lists == nil {
		return nil
	}

	result := Lists{}
	for _, list := range lists {
		if callback(list) {
			result = append(result, list)
		}
	}

	return result
}

// Strings represents a collection of strings.
type Strings []string

// Keep returns a collection containing all elements for which the given callback returns true.
func (strings Strings) Keep(callback func(string) bool) Strings {
	if strings == nil {
		return nil
	}

	result := Strings{}
	for _, s := range strings {
		if callback(s) {
			result = append(result, s)
		}
	}

	return result
}
