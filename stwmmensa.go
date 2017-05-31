package main

import "os"

// struct to store information about a dish
type dish struct {
	category string
	name     string
}

func main() {
	// parse command line arguments
	r := parseargs(os.Args)
	// get the date (today/tomorrow/monday)
	date := getDate()
	// create url
	baseURL := "http://www.studentenwerk-muenchen.de/mensa/speiseplan/"
	menuURL := baseURL + "speiseplan_" + date.Format("2006-01-02") + "_" + r.location + "_-de.html"
	// get dishes
	dishes := getDishes(menuURL)
	// create output file
	if r.format == "lis" {
		formatLIS(dishes, date, r.output)
	} else {
		formatXML(dishes, date, r.output)
	}
}
