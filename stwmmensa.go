package main

import (
	"fmt"
	"html"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// struct to store information about a dish
type dish struct {
	category string
	name     string
}

// Wochentag contains the german names of the weekdays
var Wochentag = [7]string{"Sonntag", "Montag", "Dienstag", "Mittwoch",
	"Donnerstag", "Freitag", "Samstag"}

func main() {
	// Get the date for which we want to fetch the menu.
	// TODO try to handle holidays
	now := time.Now()
	// show mondays menu during the weekend
	if now.Weekday() == time.Saturday {
		now = now.Add(time.Duration(48) * time.Hour)
	} else if now.Weekday() == time.Sunday {
		now = now.Add(time.Duration(24) * time.Hour)
	} else {
		// show next date if it is later than 14:59
		if now.Hour() > 14 {
			// for friday the next date is monday
			if now.Weekday() == time.Friday {
				now = now.Add(time.Duration(72) * time.Hour)
			} else {
				// for all other cases it is simply the next day
				now = now.Add(time.Duration(24) * time.Hour)
			}
		}
	}
	// get string representation of the date
	var date string
	date = now.Format("2006-01-02")

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
		// we are only interested in main dishes
		if strings.HasPrefix(currentDish.category, "Tagesgericht") ||
			strings.HasPrefix(currentDish.category, "Aktionsessen") ||
			strings.HasPrefix(currentDish.category, "Biogericht") ||
			currentDish.category == "Self-Service" {
			dishes = append(dishes, currentDish)
		}
	})

	outFile, err := os.Create("menu.html")
	if err != nil {
		fmt.Println(err)
	}

	titleString := "Mensa am " + Wochentag[now.Weekday()] + " den " +
		now.Format("02.01.")
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
