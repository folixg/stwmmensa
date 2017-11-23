// Package stwmmensa provides tools to get the menu for Studentenwerk Muenchen
// canteens
package stwmmensa

import (
	"encoding/xml"
	"fmt"
	"html"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Dish is used to store name and category of a dish
type Dish struct {
	Category string
	Name     string
}

// Menu consists of date, canteen location id and a list of dishes
type Menu struct {
	Date     time.Time
	Location string
	Dishes   []Dish
}

// XMLDish is used to store information about a dish in a format than can be
// used to create xml-output using xml.Marshal
type XMLDish struct {
	XMLName  xml.Name `xml:"dish"`
	Name     string   `xml:"name,attr"`
	Category string   `xml:"category,attr"`
}

// NextDayAfter is the threshold value which is used to determine, whether
// todays menu or tomorrows menu shall be fetched.
// If the processed time is equal to or greater than NextDayAfter:00 the next
// day is chosen. For example if NextDayAfter == 14 the next day will be chosen
// starting from 2PM.
const NextDayAfter = 14

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

// GermanWeekday returns the german name for a time.Weekday
func GermanWeekday(w time.Weekday) string {
	var GermanWeekdays = [7]string{"Sonntag", "Montag", "Dienstag", "Mittwoch",
		"Donnerstag", "Freitag", "Samstag"}
	return GermanWeekdays[w]
}

// GermanMonth returns the german name for a time.Month
func GermanMonth(m time.Month) string {
	var GermanMonths = [12]string{"Januar", "Februar", "März", "April", "Mai",
		"Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"}
	return GermanMonths[m-1]
}

/*
GetDate returns the date for which we want to show the canteen menu based
on the time now.
If now.Weekday() is either Saturday or Sunday the following monday will be
returned by adding 48 resp. 24 hours to now.
For monday through friday now itself will be returned if it is before
NextDayAfter PM.
If now is after NextDayAfter PM, the following day (for friday the following monday)
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
		// show next date if it is later than NextDayAfter:00
		if now.Hour() >= NextDayAfter {
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

/*
GetDishes uses goquery https://github.com/PuerkitoBio/goquery to extract
dishes from the Studentenwerk München canteen menu webpage located at url and
returns a (possibly empty) list of Dish objects.

The output should contain only main dishes. However due to the cluttered
classification by the Studentenwerk, there might still be sides/desserts which
are not handled correctly.
All declarations (e.g. (v) for "dish is vegan") are stripped from the dish names.
*/
func GetDishes(url string) []Dish {
	// fetch html from Studentenwerk
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	// extract dishes
	dishes := make([]Dish, 0, 6)
	var currentDish Dish
	doc.Find(".c-schedule__list-item").Each(func(i int, s *goquery.Selection) {
		// get category (Tagesgericht, Aktionsessen, Beilage, ...)
		// only update, if the current dish decalres a new category
		if s.Find(".stwm-artname").Text() != "" {
			currentDish.Category = s.Find(".stwm-artname").Text()
		}
		// get dish description
		currentDish.Name = s.Find(".js-schedule-dish-description").Text()
		// get rid of the "allergenkennzeichnungspflichtigen" ingredients
		currentDish.Name = strings.Split(currentDish.Name, "[")[0]
		// remove all "Zusatzstoffe" i.e. numbers enclosed in parentheses
		reZusatzstoffe, err := regexp.Compile(`\(([0-9]+,*)+\)`)
		if err != nil {
			log.Fatal(err)
		}
		currentDish.Name = reZusatzstoffe.ReplaceAllString(currentDish.Name, "")
		// remove vegetarian (f), vegan (v), beef (R), pork (S), (GQB) and (MSC)
		// declarations
		reDeclarations, err := regexp.Compile(` \(([fvRSGQBMC]+,*)+\)`)
		if err != nil {
			log.Fatal(err)
		}
		currentDish.Name = reDeclarations.ReplaceAllString(currentDish.Name, "")
		// remove any trailing whitespaces
		currentDish.Name = strings.TrimSpace(currentDish.Name)
		// we are only interested in main dishes
		// ignore everything that is of category "Beilagen"
		if currentDish.Category == "Beilagen" {
			return
		}
		// ignore everything that is of category "Beilagen"
		if currentDish.Category == "Aktion" {
			return
		}
		// ignore everything, that is just a single word
		if len(strings.Split(currentDish.Name, " ")) == 1 {
			return
		}
		// ignore known side dishes / desserts we did not catch otherwise
		if currentDish.Name == "Saisonale Beilagensalate" {
			return
		}
		if strings.HasPrefix(currentDish.Name, "Müsli mit") {
			return
		}
		if strings.HasPrefix(currentDish.Name, "Joghurt mit") {
			return
		}
		if strings.HasPrefix(currentDish.Name, "Quark mit") {
			return
		}
		dishes = append(dishes, currentDish)
	})
	return dishes
}

/*
FetchMenu fetches the menu for the canteen identified by location and the
day of date from the Studentenwerk München website.

The URL is composed in the format:
	http://www.studentenwerk-muenchen.de/mensa/speiseplan/speiseplan_2017-05-31_421_-de.html

The return value will be a (possibly empty) list of dishes.
*/
func FetchMenu(date time.Time, location string) Menu {
	var menu Menu
	menu.Date = date
	menu.Location = location
	// create url
	baseURL := "http://www.studentenwerk-muenchen.de/mensa/speiseplan/"
	menuURL := baseURL + "speiseplan_" + date.Format("2006-01-02") + "_" + location + "_-de.html"
	// get dishes
	menu.Dishes = GetDishes(menuURL)
	return menu
}

/*
FormatXML writes a XML formatted menu to the file outfile.
The menu consists of a title containing the date which was given as input
and a list of the Dish objects contained in dishes.

The output will be formatted like this:
	<menu title="Mensa am Donnerstag den 01.06.">
	<dish name="Kartoffeleintopf mit Majoran" category="Tagesgericht 1"></dish>
	<dish name="Prager Bratwurst (R,S)(2,3,8)" category="Tagesgericht 2"></dish>
	<dish name="Rindergeschnetzeltes Stroganoff (GQB) (R)(2,9)" category="Aktionsessen 5"></dish>
	<dish name="Fusilli mit Rucola-Pesto" category="Self-Service"></dish>
	<dish name="Bio-Penne mit Bio-Tomaten-Frischkäse-Sauce" category="Self-Service"></dish>
	<dish name="Kartoffeleintopf mit Majoran" category="Self-Service"></dish>
	</menu>
*/
func FormatXML(dishes []Dish, date time.Time, outfile string) {
	outFile, err := os.Create(outfile)
	if err != nil {
		fmt.Println(err)
	}
	titleString := "Mensa am " + GermanWeekday(date.Weekday()) + " den " +
		date.Format("02.01.")
	outFile.WriteString("<menu title=\"" + titleString + "\">\n")
	if len(dishes) > 0 {
		for i := range dishes {
			d := &XMLDish{Name: dishes[i].Name, Category: dishes[i].Category}
			var output []byte
			output, err = xml.Marshal(d)
			if err != nil {
				log.Fatal(err)
			}
			outFile.Write(output)
			outFile.WriteString("\n")

		}
	}
	_, err = outFile.WriteString("</menu>\n")
	if err != nil {
		log.Fatal(err)
	}
	outFile.Close()
}

/*
FormatLIS writes a HTML snippet containing the menu to the file outfile.
The menu consists of a title containing the date which was given as input
and a list of the Dish objects contained in dishes.

The output will look somethin like this:
	<h1>
	<span> class="mensa-title">Mensa</span>
	<span> class="mensa-date">Donnerstag 01. Juni 2017"</span>
	</h1>
	<div class="mensa-box">
	<div class="mensa-item">
	<span class="mensa-name">Tagesgericht 1</span>
	<span class="mensa-value">Kartoffeleintopf mit Majoran</span>
	</div>
	<div class="mensa-item">
	<span class="mensa-name">Tagesgericht 2</span>
	<span class="mensa-value">Prager Bratwurst (R,S)(2,3,8)</span>
	</div>
	<div class="mensa-item">
	<span class="mensa-name">Aktionsessen 5</span>
	<span class="mensa-value">Rindergeschnetzeltes Stroganoff (GQB) (R)(2,9)</span>
	</div>
	<div class="mensa-item">
	<span class="mensa-name">Self-Service</span>
	<span class="mensa-value">Fusilli mit Rucola-Pesto</span>
	</div>
	<div class="mensa-item">
	<span class="mensa-name">Self-Service</span>
	<span class="mensa-value">Bio-Penne mit Bio-Tomaten-Frischkäse-Sauce</span>
	</div>
	<div class="mensa-item">
	<span class="mensa-name">Self-Service</span>
	<span class="mensa-value">Kartoffeleintopf mit Majoran</span>
	</div>
	</div>
This is not valid HTML for standalone use. However this is exactly the
format needed to include the menu on the infoscreen running at the kitchen
of http://www.lis.ei.tum.de
*/
func FormatLIS(dishes []Dish, date time.Time, outfile string) {
	outFile, err := os.Create(outfile)
	if err != nil {
		fmt.Println(err)
	}
	outFile.WriteString("<h1>")
	outFile.WriteString("<span class=\"mensa-title\">Mensa</span>\n")
	titleString := GermanWeekday(date.Weekday()) + " " + date.Format("02.") + " " +
		GermanMonth(date.Month()) + date.Format(" 2006")
	// set title
	outFile.WriteString("<span class=\"mensa-date\">" +
		html.EscapeString(titleString) + "<class>\n")
	outFile.WriteString("</h1>")
	// create box
	outFile.WriteString("<div class=\"mensa-box\">\n")
	// write dishes
	if len(dishes) > 0 {
		for i := range dishes {
			outFile.WriteString("<div class=\"mensa-item\">\n")
			outFile.WriteString("<span class=\"mensa-name\">" +
				html.EscapeString(dishes[i].Category) + "</span>\n")
			outFile.WriteString("<span class=\"mensa-value\">" +
				html.EscapeString(dishes[i].Name) + "</span>\n")
			outFile.WriteString("</div>\n")
		}
	} else {
		// if we failed to fetch any dishes, assume mensa is closed
		outFile.WriteString("Leider geschlossen.")
	}
	// close box
	outFile.WriteString("</div>\n")
	outFile.Close()
}

// WriteOutput calls either FormatLIS if format=="lis" or FormatXML in all
// other cases. These methods will then write to outfile.
func WriteOutput(menu Menu, format string, outfile string) {
	if format == "lis" {
		FormatLIS(menu.Dishes, menu.Date, outfile)
	} else {
		FormatXML(menu.Dishes, menu.Date, outfile)
	}
}

// UpdateMenuFile updates (or creates) the file "outfile" with the current
// menu for canteen with id "location" formatted in the "format" style.
func UpdateMenuFile(location string, format string, outfile string) {
	// get the date (today/tomorrow/monday)
	date := GetDate(time.Now())
	// fetch the menu
	menu := FetchMenu(date, location)
	// create output file
	WriteOutput(menu, format, outfile)
}
