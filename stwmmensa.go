package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// basic struct to store information about a dish
type dish struct {
	category string
	name     string
}

// xml representation of a dish
type xmlDish struct {
	XMLName  xml.Name `xml:"dish"`
	Name     string   `xml:"name,attr"`
	Category string   `xml:"category,attr"`
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

	outFile, err := os.Create("menu.xml")

	if err != nil {
		fmt.Println(err)
	}
	_, err = outFile.WriteString(xml.Header)
	if err != nil {
		log.Fatal(err)
	}

	var titleString string
	// if we failed to fetch any dishes, assume mensa is closed
	if len(dishes) == 0 {
		titleString = "Leider hat die Mensa am " + Wochentag[now.Weekday()] + " den " +
			now.Format("02.01.") + " geschlossen."
	} else {
		titleString = "Am " + Wochentag[now.Weekday()] + " den " +
			now.Format("02.01.") + " empfiehlt Küchenchef Horst Waldner:"
	}
	_, err = outFile.WriteString("<menu title=\"" + titleString + "\">\n")
	if err != nil {
		log.Fatal(err)
	}
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
