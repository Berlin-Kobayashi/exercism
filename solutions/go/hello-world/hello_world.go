/*
Package greeting implements a simple library for greeting.
*/
package greeting

import (
	"fmt"
)

const testVersion = 3

// HelloWorld returns a greeting for the given name. If name is an empty string the world is greeted.
func HelloWorld(name string) string {
	greetingFormat := "Hello, %s!"

	if name == "" {
		name = "World"
	}

	return fmt.Sprintf(greetingFormat, name)
}
