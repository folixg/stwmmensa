package main

import (
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

	// create output file
	if r.format == "lis" {
		formatLIS(dishes, date, r.output)
	} else {
		formatXML(dishes, date, r.output)
	}
}
