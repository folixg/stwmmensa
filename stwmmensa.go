package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// parse command line arguments
	r, err := parseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	// get the date (today/tomorrow/monday)
	date := getDate(time.Now())
	// fetch the menu
	menu := fetchMenu(date, r.location)

	// create output file
	writeOutput(menu, r.format, r.output)
}
