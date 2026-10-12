package clock

import (
	"fmt"
	"time"
)

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func getTime(h, m int) time.Time {
	year := 2024
	month := time.July
	day := 8
	hour := h
	minute := m
	second := 0
	nanosecond := 0
	location := time.Now().Location() // Use the local time zone

	return time.Date(year, month, day, hour, minute, second, nanosecond, location)
}

func New(h, m int) Clock {
	specificTime := getTime(h, m)
	return Clock{specificTime.Hour(), specificTime.Minute()}
}

func (c Clock) Add(m int) Clock {
	duration := time.Duration(m) * time.Minute
	specificTime := getTime(c.hour, c.minute)
	res := specificTime.Add(duration)

	return Clock{res.Hour(), res.Minute()}
}

func (c Clock) Subtract(m int) Clock {
	duration := -time.Duration(m) * time.Minute
	specificTime := getTime(c.hour, c.minute)
	res := specificTime.Add(duration)

	return Clock{res.Hour(), res.Minute()}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
