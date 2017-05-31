package main

import (
	"time"
)

func getDate() time.Time {
	// Get the date for which we want to fetch the menu.
	// TODO try to handle holidays
	now := time.Now()
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

func wochentag(w time.Weekday) string {
	var Wochentag = [7]string{"Sonntag", "Montag", "Dienstag", "Mittwoch",
		"Donnerstag", "Freitag", "Samstag"}
	return Wochentag[w]
}
