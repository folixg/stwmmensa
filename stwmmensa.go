package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the date for which we want to fetch the menu.
	// Which is today until 15:00, and tomorrow after that
	// TODO handle weekends and holidays
	now := time.Now()
	var date string
	if now.Hour() < 23 { // TODO change back to 15
		date = now.Format("2006-01-02")
	} else {
		date = now.Add(time.Duration(9) * time.Hour).Format("2006-01-02")
	}
	fmt.Println(date)
	// TODO different locations
	location := "421"
	// create url
	baseURL := "http://www.studentenwerk-muenchen.de/mensa/speiseplan/"
	menuPage := baseURL + "speiseplan_" + date + "_" + location + "_-de.html"
	fmt.Println(menuPage)
	// TODO fetch html
	// TODO parse html
	// TODO output

}
