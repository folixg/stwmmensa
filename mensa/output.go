package mensa

import (
	"encoding/xml"
	"fmt"
	"html"
	"log"
	"os"
	"time"
)

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
	<h1 class="mensa-title">Mensa am Donnerstag den 01.06.</h1>
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
	titleString := "Mensa am " + GermanWeekday(date.Weekday()) + " den " +
		date.Format("02.01.")
	// set title
	outFile.WriteString("<h1 class=\"mensa-title\">" +
		html.EscapeString(titleString) + "</h1>\n")
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
