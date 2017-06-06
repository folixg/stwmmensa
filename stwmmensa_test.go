package stwmmensa

import (
	"testing"
	"time"
)

func TestLocationValid(t *testing.T) {
	cases := []struct {
		in       string
		expected bool
	}{
		{"411", true},
		{"412", true},
		{"421", true},
		{"422", true},
		{"423", true},
		{"431", true},
		{"432", true},
		{"420", false},
		{"xxx", false},
	}
	for _, c := range cases {
		got := LocationValid(c.in)
		if got != c.expected {
			t.Errorf("LocationValid(%q) == %t, expected %t\n", c.in, got, c.expected)
		}
	}
}

func TestFormatValid(t *testing.T) {
	cases := []struct {
		in       string
		expected bool
	}{
		{"xml", true},
		{"lis", true},
		{"html", false},
	}
	for _, c := range cases {
		got := FormatValid(c.in)
		if got != c.expected {
			t.Errorf("FormatValid(%q) == %t, expected %t\n", c.in, got, c.expected)
		}
	}
}

func TestGermanWeekday(t *testing.T) {
	cases := []struct {
		in       time.Weekday
		expected string
	}{
		{0, "Sonntag"},
		{1, "Montag"},
		{2, "Dienstag"},
		{3, "Mittwoch"},
		{4, "Donnerstag"},
		{5, "Freitag"},
		{6, "Samstag"},
	}
	for _, c := range cases {
		got := GermanWeekday(c.in)
		if got != c.expected {
			t.Errorf("GermanWeekday(%d) == %q, expected %q\n", c.in, got, c.expected)
		}
	}
}

func TestGetDate(t *testing.T) {
	cases := []struct {
		in       time.Time
		expected time.Time
	}{
		// normal weekday before NextDayAfter
		{time.Date(2017, 05, 31, 0, 0, 0, 0, time.Local),
			time.Date(2017, 05, 31, 0, 0, 0, 0, time.Local)},
		// normal weekday after NextDayAfter
		{time.Date(2017, 05, 31, 16, 0, 0, 0, time.Local),
			time.Date(2017, 06, 01, 16, 0, 0, 0, time.Local)},
		// friday before NextDayAfter
		{time.Date(2017, 05, 26, 0, 0, 0, 0, time.Local),
			time.Date(2017, 05, 26, 0, 0, 0, 0, time.Local)},
		// friday after NextDayAfter
		{time.Date(2017, 05, 26, 16, 0, 0, 0, time.Local),
			time.Date(2017, 05, 29, 16, 0, 0, 0, time.Local)},
		// saturday before NextDayAfter
		{time.Date(2017, 05, 27, 0, 0, 0, 0, time.Local),
			time.Date(2017, 05, 29, 0, 0, 0, 0, time.Local)},
		// saturday after NextDayAfter
		{time.Date(2017, 05, 27, 16, 0, 0, 0, time.Local),
			time.Date(2017, 05, 29, 16, 0, 0, 0, time.Local)},
		// sunday before NextDayAfter
		{time.Date(2017, 06, 4, 0, 0, 0, 0, time.Local),
			time.Date(2017, 06, 5, 0, 0, 0, 0, time.Local)},
		// sunday after NextDayAfter
		{time.Date(2017, 06, 4, 16, 0, 0, 0, time.Local),
			time.Date(2017, 06, 5, 16, 0, 0, 0, time.Local)},
	}
	for _, c := range cases {
		got := GetDate(c.in)
		if got != c.expected {
			t.Errorf("GetDate(%s) == %s, expected %s\n", c.in, got, c.expected)
		}
	}
}

func TestGetDishes(t *testing.T) {
	expected := []Dish{
		{Category: "Tagesgericht 1",
			Name: "Zartweizen mit Tomaten, Zucchini und Auberginen"},
		{Category: "Tagesgericht 3",
			Name: "Cevapcici von der Pute mit Ajvar"},
		{Category: "Aktionsessen 4",
			Name: "Münchner Biergulasch (GQB) (R)(99)"},
		{Category: "Self-Service",
			Name: "Zartweizen mit Tomaten, Zucchini und Auberginen"},
		{Category: "Self-Service",
			Name: "Rigatoni mit Paprikapesto"},
		{Category: "Self-Service",
			Name: "Bunte Nudel-Hackfleisch-Pfanne (R)"},
	}
	got := GetDishes("http://www.studentenwerk-muenchen.de/mensa/speiseplan/speiseplan_2017-05-22_421_-de.html")
	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("Expected: %s, %s\n Got: %s, %s", expected[i].Category,
				expected[i].Name, got[i].Category, got[i].Name)
		}
	}
}
