package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type args struct {
	location string
	output   string
	format   string
}

func printHelp() {
	fmt.Print(
		`stwmmensa accepts the following arguments:
-h --help: Print this usage info and quit
-o --output: set the path to the output file (mandatory)
-l --location: set mensa location id (default: 421)
-f --format: select format (default: xml)

** Location
Here is a list of valid location codes with the name of the corresponding mensa:
411 : Mensa Leopoldstraße
412 : Mensa Martinsried
421 : Mensa Arcisstraße
422 : Mensa Garching
423 : Mensa Weihenstephan
431 : Mensa Lothstraße
432 : Mensa Pasing

** Format
xml : a generic xml file is created
lis : generate html snippet for LIS-infoscreen

** Examples:
stwmmensa -l 411 -o /my/path/leopold.xml -f xml
stwmmensa --location=423 --output=weihenstephan.html --format=lis
`)
	os.Exit(0)
}

func locationValid(id string) bool {
	var mensen = map[string]string{
		"411": "Mensa Leopoldstraße",
		"412": "Mensa Martinsried",
		"421": "Mensa Arcisstraße",
		"422": "Mensa Garching",
		"423": "Mensa Weihenstephan",
		"431": "Mensa Lothstraße",
		"432": "Mensa Pasing",
	}
	var valid bool
	_, valid = mensen[id]
	return valid
}

func formatValid(format string) bool {
	return format == "xml" || format == "lis"
}

func parseArgs(osargs []string) args {
	// initialize internal arguents with empty strings
	var r args
	r.location = ""
	r.output = ""
	r.format = ""

	i := 1
	for i < len(osargs) {
		switch {
		case osargs[i] == "-h", osargs[i] == "--help":
			printHelp()
		case osargs[i] == "-l":
			r.location = osargs[i+1]
			i += 2
		case strings.HasPrefix(osargs[i], "--location="):
			r.location = strings.TrimPrefix(osargs[i], "--location=")
			i += 1
		case osargs[i] == "-o":
			r.output = osargs[i+1]
			i += 2
		case strings.HasPrefix(osargs[i], "--output="):
			r.output = strings.TrimPrefix(osargs[i], "--output=")
			i += 1
		case osargs[i] == "-f":
			r.format = osargs[i+1]
			i += 2
		case strings.HasPrefix(osargs[i], "--format="):
			r.format = strings.TrimPrefix(osargs[i], "--format=")
			i += 1
		default:
			log.Fatal("Unknown argument " + osargs[i] + ". Run stwmmensa -h for help.")
		}
	}

	// was the mandatory output file provided?
	if r.output == "" {
		log.Fatal("No output file provided. Run stwmmensa -h for help.")
	}
	// set default values for argmunets we did not get via command line
	if r.location == "" {
		r.location = "421"
	}
	if r.format == "" {
		r.format = "xml"
	}
	// check if we have valid values
	if !locationValid(r.location) {
		log.Fatal(r.location + " is not a valid location identifier. Run stwmmensa -h for help.")
	}
	if !formatValid(r.format) {
		log.Fatal(r.format + " is not a valid format. Run stwmmensa -h for help.")
	}

	return r
}
