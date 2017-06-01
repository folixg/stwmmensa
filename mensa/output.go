package main

import (
	"encoding/xml"
	"fmt"
	"html"
	"log"
	"os"
	"time"
)

// xml representation of a dish
type xmlDish struct {
	XMLName  xml.Name `xml:"dish"`
	Name     string   `xml:"name,attr"`
	Category string   `xml:"category,attr"`
}

func formatXML(dishes []dish, date time.Time, outfile string) {
	outFile, err := os.Create(outfile)
	if err != nil {
		fmt.Println(err)
	}
	titleString := "Mensa am " + wochentag(date.Weekday()) + " den " +
		date.Format("02.01.")
	outFile.WriteString("<menu title=\"" + titleString + "\">\n")
	if len(dishes) > 0 {
		for i := range dishes {
			d := &xmlDish{Name: dishes[i].name, Category: dishes[i].category}
			output, err := xml.Marshal(d)
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

func formatLIS(dishes []dish, date time.Time, outfile string) {
	outFile, err := os.Create(outfile)
	if err != nil {
		fmt.Println(err)
	}
	titleString := "Mensa am " + wochentag(date.Weekday()) + " den " +
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
				html.EscapeString(dishes[i].category) + "</span>\n")
			outFile.WriteString("<span class=\"mensa-value\">" +
				html.EscapeString(dishes[i].name) + "</span>\n")
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

func writeOutput(menu Menu, format string, outfile string) {
	if format == "lis" {
		formatLIS(menu.dishes, menu.date, outfile)
	} else {
		formatXML(menu.dishes, menu.date, outfile)
	}
}
