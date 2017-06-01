package mensa

import (
	"time"
)

/*
GetDate returns the date for which we want to show the canteen menu based
on the time now.
If now.Weekday() is either Saturday or Sunday the following monday will be
returned by adding 48 resp. 24 hours to now.
For monday through friday now itself will be returned if it is before 3PM.
If now is after 3PM, the following day (for friday the following monday)
will be returned by adding 24 hours (72 hours for friday) to now.
*/
func GetDate(now time.Time) time.Time {
	// BUG(tg) there is no handling of holidays

	// show mondays menu during the weekend
	if now.Weekday() == time.Saturday {
		now = now.Add(time.Duration(48) * time.Hour)
	} else if now.Weekday() == time.Sunday {
		now = now.Add(time.Duration(24) * time.Hour)
	} else {
		// show next date if it is later than 14:59
		if now.Hour() > 14 {
			// for friday the next date is monday
			if now.Weekday() == time.Friday {
				now = now.Add(time.Duration(72) * time.Hour)
			} else {
				// for all other cases it is simply the next day
				now = now.Add(time.Duration(24) * time.Hour)
			}
		}
	}
	return now
}

// GermanWeekday returns the german name for a time.Weekday
func GermanWeekday(w time.Weekday) string {
	var GermanWeekdays = [7]string{"Sonntag", "Montag", "Dienstag", "Mittwoch",
		"Donnerstag", "Freitag", "Samstag"}
	return GermanWeekdays[w]
}
