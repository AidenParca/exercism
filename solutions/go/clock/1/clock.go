package clock

import "fmt"

// minutesInDay represents the total number of minutes in a 24-hour day.
const minutesInDay = 24 * 60
// minutesInHour represents the number of minutes in an hour.
const minutesInHour = 60

// Clock represents a time of day in minutes from midnight.
// This integer value always stays within the range of 0 to 1439 (minutesInDay - 1).
// Using a simple integer makes comparisons (==) and arithmetic straightforward.
type Clock int

// New creates a new Clock instance.
// It correctly handles hour and minute rollovers for both positive and negative values,
// normalizing the time to a 24-hour format.
func New(h, m int) Clock {
	// Calculate the total number of minutes from midnight.
	totalMinutes := h*minutesInHour + m

	// Normalize the minutes to be within a single day [0, 1439].
	// The formula `(a % n + n) % n` is a standard way to ensure the
	// result of a modulo operation is always non-negative.
	normalizedMinutes := (totalMinutes%minutesInDay + minutesInDay) % minutesInDay

	return Clock(normalizedMinutes)
}

// Add returns a new Clock with the given number of minutes added.
// It handles rollovers into the next or previous day(s).
func (c Clock) Add(m int) Clock {
	// To add minutes, we create a new clock from the sum of the current
	// minutes and the minutes to add. The New function handles the normalization.
	return New(0, int(c)+m)
}

// Subtract returns a new Clock with the given number of minutes subtracted.
// It handles rollovers into the next or previous day(s).
func (c Clock) Subtract(m int) Clock {
	// Subtracting minutes is equivalent to adding a negative number of minutes.
	return c.Add(-m)
}

// String returns a string representation of the clock in "HH:MM" format.
func (c Clock) String() string {
	// Calculate the hour and minute from the total minutes stored in the Clock.
	hour := int(c) / minutesInHour
	minute := int(c) % minutesInHour

	// Use fmt.Sprintf for idiomatic zero-padded formatting.
	return fmt.Sprintf("%02d:%02d", hour, minute)
}