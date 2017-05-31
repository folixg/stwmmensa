package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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
		// we are only interested in main dishes
		if strings.HasPrefix(currentDish.category, "Tagesgericht") ||
			strings.HasPrefix(currentDish.category, "Aktionsessen") ||
			strings.HasPrefix(currentDish.category, "Biogericht") ||
			currentDish.category == "Self-Service" {
			dishes = append(dishes, currentDish)
		}
	})
	return dishes
}
