package booking

import "time"
import "fmt"
// Schedule returns a time.Time from a string containing a date.
// "7/25/2019 13:45:00"
// formattedTime


func Schedule(date string) time.Time {
    layout := "1/02/2006 15:04:05"
	p,_ := time.Parse(layout,date)
    return p // time.Time, error
	panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    parsedTime, err := time.Parse(layout, date)
    if err != nil {
        return false
    }
    now := time.Now()
    return !parsedTime.After(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    // 		   Friday, March 8, 1974 12:02:02
    layout := "Monday, January 2, 2006 15:04:05"
    parsedTime, _ := time.Parse(layout, date)
    if parsedTime.Hour() > 11 && parsedTime.Hour() < 19{
        return true
    } else {
        return false
    }
    
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	
	t, _ := time.Parse(layout, date)

	dateLayout := "Monday, January 2, 2006"
	timeLayout := "15:04"

	return fmt.Sprintf(
		"You have an appointment on %s, at %s.",
		t.Format(dateLayout),
		t.Format(timeLayout),
	)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	panic("Please implement the AnniversaryDate function")
}
