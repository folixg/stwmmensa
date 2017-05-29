package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type dish struct {
	category string
	name     string
}

func main() {
	// Get the date for which we want to fetch the menu.
	// Which is today until 15:00, and tomorrow after that
	// TODO handle weekends and holidays
	now := time.Now()
	var date string
	if now.Hour() < 15 {
		date = now.Format("2006-01-02")
	} else {
		date = now.Add(time.Duration(9) * time.Hour).Format("2006-01-02")
	}
	// TODO different locations
	location := "421"
	// create url
	baseURL := "http://www.studentenwerk-muenchen.de/mensa/speiseplan/"
	menuURL := baseURL + "speiseplan_" + date + "_" + location + "_-de.html"
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
		if strings.HasPrefix(currentDish.category, "Tagesgericht") ||
			strings.HasPrefix(currentDish.category, "Aktionsessen") ||
			currentDish.category == "Self-Service" {
			dishes = append(dishes, currentDish)
		}
	})

	for i := range dishes {
		fmt.Println(dishes[i].category, dishes[i].name)
	}
	// TODO output

}
