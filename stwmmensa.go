package main

import (
	"fmt"
	"html"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// struct to store information about a dish
type dish struct {
	category string
	name     string
}

func main() {
	// parse command line arguments
	r := parseargs(os.Args)

	// get string representation of the date
	date := getDate()

	// create url
	baseURL := "http://www.studentenwerk-muenchen.de/mensa/speiseplan/"
	menuURL := baseURL + "speiseplan_" + date.Format("2006-01-02") + "_" + r.location + "_-de.html"
	// fetch html from Studentenwerk
	doc, err := goquery.NewDocument(menuURL)
	if err != nil {
		log.Fatal(err)
	}
	// extract dishes
	dishes := make([]dish, 0, 6)
	var currentDish dish
	doc.Find(".c-schedule__list-item").Each(func(i int, s *goquery.Selection) {
		// get category (Tagesgericht, Aktionsessen, Beilage, ...)
		currentDish.category = s.Find(".stwm-artname").Text()
		// get dish description
		currentDish.name = s.Find(".js-schedule-dish-description").Text()
		// get rid of the "allergenkennzeichnungspflichtigen" ingredients
		currentDish.name = strings.Split(currentDish.name, "[")[0]
		// we are only interested in main dishes
		if strings.HasPrefix(currentDish.category, "Tagesgericht") ||
			strings.HasPrefix(currentDish.category, "Aktionsessen") ||
			strings.HasPrefix(currentDish.category, "Biogericht") ||
			currentDish.category == "Self-Service" {
			dishes = append(dishes, currentDish)
		}
	})

	outFile, err := os.Create(r.output)
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
