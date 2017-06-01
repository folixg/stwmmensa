// Package mensa provides methods and types used to fetch canteen menus
// from the Studentenwerk München website.
package mensa

import (
	"encoding/xml"
	"time"
)

// Args is used to store command line arguments.
type Args struct {
	Location string
	Outfile  string
	Format   string
}

// Dish is used to store name and category of a dish
type Dish struct {
	Category string
	Name     string
}

// Menu consists of date, canteen location id and a list of dishes
type Menu struct {
	Date     time.Time
	Location string
	Dishes   []Dish
}

// XMLDish is used to store information about a dish in a format than can be
// used to create xml-output using xml.Marshal
type XMLDish struct {
	XMLName  xml.Name `xml:"dish"`
	Name     string   `xml:"name,attr"`
	Category string   `xml:"category,attr"`
}
