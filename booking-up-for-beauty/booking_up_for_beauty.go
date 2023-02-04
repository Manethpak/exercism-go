package booking

import (
	"strconv"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	
	t, err := time.Parse("January 2, 2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if (err != nil) {
		panic(err)
	}
	hour := t.Hour()
	if (hour >= 12 && hour < 18) {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	result := t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	return result
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	t, err := time.Parse("1/2/2006", "9/15/" + strconv.Itoa(year))
	if err != nil {
		panic(err)
	}
	return t
}
