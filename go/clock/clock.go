package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	minutes := m % 60

	if minutes < 0 {
		minutes += 60
		h--
	}

	hours := (h + (m / 60)) % 24

	if hours < 0 {
		hours += 24
	}

	return Clock{hours, minutes}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, (c.minute + m))
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, (c.minute - m))
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
