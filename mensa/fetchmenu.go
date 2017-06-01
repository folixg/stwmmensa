package main

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// struct to store information about a dish
type dish struct {
	category string
	name     string
}

// menu consists of a date, a location and a slice of dishes
type Menu struct {
	date     time.Time
	location string
	dishes   []dish
}

func getDishes(url string) []dish {
	// fetch html from Studentenwerk
	doc, err := goquery.NewDocument(url)
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
		currentDish.name = strings.Trim(currentDish.name, " ")
		// we are only interested in main dishes
		// TODO add support for "Mensa Spezial", "Mensa Klassiker", ...
		if strings.HasPrefix(currentDish.category, "Tagesgericht") ||
			strings.HasPrefix(currentDish.category, "Aktionsessen") ||
			strings.HasPrefix(currentDish.category, "Biogericht") ||
			currentDish.category == "Self-Service" {
			dishes = append(dishes, currentDish)
		}
	})
	return dishes
}

func fetchMenu(date time.Time, location string) Menu {
	var menu = new(Menu)
	menu.date = date
	menu.location = location
	// create url
	baseURL := "http://www.studentenwerk-muenchen.de/mensa/speiseplan/"
	menuURL := baseURL + "speiseplan_" + date.Format("2006-01-02") + "_" + location + "_-de.html"
	// get dishes
	menu.dishes = getDishes(menuURL)
	return *menu
}
