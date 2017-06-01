package mensa

import (
	"encoding/xml"
	"time"
)

type Args struct {
	Location string
	Outfile  string
	Format   string
}

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

// xml representation of a dish
type xmlDish struct {
	XMLName  xml.Name `xml:"dish"`
	Name     string   `xml:"name,attr"`
	Category string   `xml:"category,attr"`
}
