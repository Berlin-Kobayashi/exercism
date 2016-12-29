// Package gigasecond implements a function for adding a gigasecond to a Time.
package gigasecond

import (
	"time"
)

const testVersion = 4

// AddGigasecond adds a gigasecond (one billion seconds) to the given Time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1e9) * time.Second)
}
