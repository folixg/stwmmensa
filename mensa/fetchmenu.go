package mensa

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

/*
GetDishes uses goquery https://github.com/PuerkitoBio/goquery to extract
dishes from the Studentenwerk München canteen menu webpage located at url and
returns a (possibly empty) list of Dish objects.

At the moment only main dishes are returned.
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
		currentDish.Category = s.Find(".stwm-artname").Text()
		// get dish description
		currentDish.Name = s.Find(".js-schedule-dish-description").Text()
		// get rid of the "allergenkennzeichnungspflichtigen" ingredients
		currentDish.Name = strings.Split(currentDish.Name, "[")[0]
		currentDish.Name = strings.Trim(currentDish.Name, " ")
		// we are only interested in main dishes
		if strings.HasPrefix(currentDish.Category, "Tagesgericht") ||
			strings.HasPrefix(currentDish.Category, "Aktionsessen") ||
			strings.HasPrefix(currentDish.Category, "Biogericht") ||
			currentDish.Category == "Self-Service" {
			dishes = append(dishes, currentDish)
			// BUG(tg) canteens which use a naming scheme different from "Tagesgericht",
			// "Aktionsessen", "Biogericht" and "Self-Service" won't work at the moment
		}
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
