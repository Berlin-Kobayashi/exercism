// Package meetup implements a library for calculating the date of meetups.
package meetup

import "time"

const testVersion = 3

// WeekSchedule represents the week schedule in which a monthly meetup happens.
type WeekSchedule int

const (
	// First represents the week schedule in which a meetup happens on the first occurence of a weekday in a month.
	First WeekSchedule = iota
	// Second represents the week schedule in which a meetup happens on the second occurence of a weekday in a month.
	Second
	// Third represents the week schedule in which a meetup happens on the third occurence of a weekday in a month.
	Third
	// Fourth represents the week schedule in which a meetup happens on the fourth occurence of a weekday in a month.
	Fourth
	// Last represents the week schedule in which a meetup happens on the last occurence of a weekday in a month.
	Last
	// Teenth represents the week schedule in which a meetup happens on the first occurence of a weekday in the -teenth days of a month.
	Teenth
)

// Day returns the day of the month in which the meetup with the given schedule will happen.
func Day(weekSchedule WeekSchedule, weekday time.Weekday, month time.Month, year int) (dayOfMonth int) {
	var weekStartingDate time.Time
	switch weekSchedule {
	case First, Second, Third, Fourth:
		weekStartingDate = time.Date(year, month, int(weekSchedule)*7+1, 0, 0, 0, 0, time.UTC)
	case Last:
		firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		weekStartingDate = firstDayOfMonth.AddDate(0, 1, -7)
	case Teenth:
		weekStartingDate = time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
	default:
		return 0
	}

	weekStartingWeekday := weekStartingDate.Weekday()
	if weekday >= weekStartingWeekday {
		return weekStartingDate.Day() + int(weekday) - int(weekStartingWeekday)
	}

	return weekStartingDate.Day() + 7 - (int(weekStartingWeekday) - int(weekday))
}
