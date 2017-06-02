/*
stwmmensa fetches the canteen menu from the Studentenwerk Muenchen website

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
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/folixg/stwmmensa/mensa"
)

const help = `stwmmensa accepts the following arguments:
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

`

func main() {
	// parse command line arguments
	args, err := mensa.ParseArgs(os.Args)
	if err != nil {
		if err.Error() == "help" {
			fmt.Print(help)
			os.Exit(0)
		}
		log.Fatal(err)
	}
	// get the date (today/tomorrow/monday)
	date := mensa.GetDate(time.Now())
	// fetch the menu
	menu := mensa.FetchMenu(date, args.Location)
	// create output file
	mensa.WriteOutput(menu, args.Format, args.Outfile)
}
