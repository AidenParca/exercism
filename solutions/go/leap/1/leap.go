package leap

// IsLeapYear determines if the year given leaped in the calendar or not
func IsLeapYear(year int) bool {
	if year % 4 != 0 {
        return false
    }
    if year % 4 == 0 {
        switch  {
        case year % 100 == 0 : {if year % 400 == 0 {
            return true
        } else {
            return false
        }}
        case year % 100 != 0 : return true
    	}
        return true
    }
    return false
    
}
