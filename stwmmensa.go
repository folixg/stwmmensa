package main

import (
	"fmt"
	"log"
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
	if now.Hour() < 23 { // TODO change back to 15
		date = now.Format("2006-01-02")
	} else {
		date = now.Add(time.Duration(9) * time.Hour).Format("2006-01-02")
	}
	// TODO different locations
	location := "421"
	// create url
	baseURL := "http://www.studentenwerk-muenchen.de/mensa/speiseplan/"
	menuURL := baseURL + "speiseplan_" + date + "_" + location + "_-de.html"
	// TODO fetch html
	doc, err := goquery.NewDocument(menuURL)
	if err != nil {
		log.Fatal(err)
	}

	dishes := make([]dish, 20) // TODO get actual size needed

	doc.Find(".c-schedule__list-item").Each(func(i int, s *goquery.Selection) {
		dishes[i].category = s.Find(".stwm-artname").Text()
		if dishes[i].category == "" {
			dishes[i].category = dishes[i-1].category
		}
		dishes[i].name = s.Find(".js-schedule-dish-description").Text()
	})

	for i := range dishes {
		fmt.Println(dishes[i].category, dishes[i].name)
	}
	// TODO output

}
