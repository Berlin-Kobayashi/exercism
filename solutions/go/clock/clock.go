/*
Package clock implements a simple clock type which represents a moment by hours and minutes.
*/
package clock

import (
	"fmt"
)

const testVersion = 4

// Clock represents a moment by hours and minutes.
type Clock struct {
	Hour, Minute int
}

// New returns a new Clock.
func New(hour, minute int) Clock {
	clockHour := (hour + minute/60) % 24
	clockMinute := minute % 60

	if clockMinute < 0 {
		clockMinute += 60
		clockHour--
	}
	if clockHour < 0 {
		clockHour += 24
	}

	return Clock{clockHour, clockMinute}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.Hour, clock.Minute)
}

// Add returns a new Clock with added minutes. A negative value can be passed to substract minutes.
func (clock Clock) Add(minutes int) Clock {
	return New(clock.Hour, clock.Minute+minutes)
}
