package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	ti, _ := time.Parse(layout, date)

	return ti
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// "July 25, 2019 13:45:00"
	layout := "January 2, 2006 15:04:05"
	ti, _ := time.Parse(layout, date)

	return ti.Unix() <= time.Now().Unix()
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// "Thursday, July 25, 2019 13:45:00"
	layout := "Monday, January 2, 2006 15:04:05"

	ti, _ := time.Parse(layout, date)

	return ti.Hour() >= 12 && ti.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// "7/25/2019 13:45:00"
	layout := "1/2/2006 15:04:05"
	ti, _ := time.Parse(layout, date)

	// Thursday, July 25, 2019, at 13:45."
	formattedTime := ti.Format("Monday, January 2, 2006, at 15:04")

	return fmt.Sprintf("You have an appointment on %s.", formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02"

	year := time.Now().Year()
	ti, _ := time.Parse(layout, fmt.Sprintf("%d-09-15", year))

	return ti
}
