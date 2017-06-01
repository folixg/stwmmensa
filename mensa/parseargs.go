package mensa

import (
	"fmt"
	"strings"
)

/*
PrintHelp prints following text to stdout:
	stwmmensa accepts the following arguments:
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
*/
func PrintHelp() {
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
}

// LocationValid checks whether string id is a valid STWM canteen ID
func LocationValid(id string) bool {
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

// FormatValid checks, whether the string format is either "xml" or "lis",
// which are the two supported output formats
func FormatValid(format string) bool {
	return format == "xml" || format == "lis"
}

/*
ParseArgs checks if the command line arguments are valid.
The input osargs is supposed to be the output of os.Args

Known arguments are:
	-h --help: Print this usage info and quit
	-o --output: set the path to the output file (mandatory)
	-l --location: set mensa location id (default: 421)
	-f --format: select format (default: xml)

If -h or --help is detected, a help message is printed to stdout.

If the parsed command line aguments are either not valid or contained the
-h/--help switch, a non-nil error is returned and the returned arguments
Args do not contain valid command line arguments.
*/
func ParseArgs(osargs []string) (Args, error) {
	// initialize internal arguents with empty strings
	var r Args
	r.Location = ""
	r.Outfile = ""
	r.Format = ""

	i := 1
	for i < len(osargs) {
		switch {
		case osargs[i] == "-h", osargs[i] == "--help":
			//printHelp()
			return r, fmt.Errorf("print help")
		case osargs[i] == "-l":
			r.Location = osargs[i+1]
			i += 2
		case strings.HasPrefix(osargs[i], "--location="):
			r.Location = strings.TrimPrefix(osargs[i], "--location=")
			i++
		case osargs[i] == "-o":
			r.Outfile = osargs[i+1]
			i += 2
		case strings.HasPrefix(osargs[i], "--output="):
			r.Outfile = strings.TrimPrefix(osargs[i], "--output=")
			i++
		case osargs[i] == "-f":
			r.Format = osargs[i+1]
			i += 2
		case strings.HasPrefix(osargs[i], "--format="):
			r.Format = strings.TrimPrefix(osargs[i], "--format=")
			i++
		default:
			return r, fmt.Errorf("unknown argument " + osargs[i] + ". Run stwmmensa -h for help")
		}
	}

	// was the mandatory output file provided?
	if r.Outfile == "" {
		return r, fmt.Errorf("no output file provided. Run stwmmensa -h for help")
	}
	// set default values for argmunets we did not get via command line
	if r.Location == "" {
		r.Location = "421"
	}
	if r.Format == "" {
		r.Format = "xml"
	}
	// check if we have valid values
	if !LocationValid(r.Location) {
		return r, fmt.Errorf(r.Location + " is not a valid location identifier. Run stwmmensa -h for help")
	}
	if !FormatValid(r.Format) {
		return r, fmt.Errorf(r.Format + " is not a valid format. Run stwmmensa -h for help")
	}

	return r, nil
}
