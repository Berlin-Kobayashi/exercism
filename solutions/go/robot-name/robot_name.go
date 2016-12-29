// Package robotname implements library for generating random robot names with minimal collision.
package robotname

import (
	"math/rand"
)

// Robot represents a nameable robot.
type Robot struct {
	name string
}

var names = map[string]bool{}

const maxUniqueNames = 26 * 26 * 10 * 10 * 10

// Name gives the robot a random name if it is yet unnamed and returns it.
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = generateRandomName()
		if !areAllNamesUsed() {
			for names[r.name] {
				r.name = generateRandomName()
			}

			names[r.name] = true
		}

	}

	return r.name
}

func generateRandomName() string {
	randomNumber := rand.Intn(maxUniqueNames + 1)
	result := make([]byte, 5, 5)
	result[4] = byte('0' + randomNumber%10)
	randomNumber /= 10
	result[3] = byte('0' + randomNumber%10)
	randomNumber /= 10
	result[2] = byte('0' + randomNumber%10)
	randomNumber /= 10
	result[1] = byte('A' + randomNumber%26)
	randomNumber /= 26
	result[0] = byte('A' + randomNumber%26)
	randomNumber /= 26

	return string(result)
}

func areAllNamesUsed() bool {
	return len(names) == maxUniqueNames
}

// Reset resets the name of the robot.
func (r *Robot) Reset() {
	r.name = ""
}
