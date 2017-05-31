package main

import (
	"testing"
	"time"
)

func TestWochentag(t *testing.T) {
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
		got := wochentag(c.in)
		if got != c.expected {
			t.Errorf("wochentag(%d) == %q, expected %q\n", c.in, got, c.expected)
		}
	}
}

func TestDate(t *testing.T) {
	cases := []struct {
		in       time.Time
		expected time.Time
	}{
		// normal weekday before 15:00
		{time.Date(2017, 05, 31, 0, 0, 0, 0, time.Local),
			time.Date(2017, 05, 31, 0, 0, 0, 0, time.Local)},
		// normal weekday after 15:00
		{time.Date(2017, 05, 31, 16, 0, 0, 0, time.Local),
			time.Date(2017, 06, 01, 16, 0, 0, 0, time.Local)},
		// friday before 15:00
		{time.Date(2017, 05, 26, 0, 0, 0, 0, time.Local),
			time.Date(2017, 05, 26, 0, 0, 0, 0, time.Local)},
		// friday after 15:00
		{time.Date(2017, 05, 26, 16, 0, 0, 0, time.Local),
			time.Date(2017, 05, 29, 16, 0, 0, 0, time.Local)},
		// saturday before 15:00
		{time.Date(2017, 05, 27, 0, 0, 0, 0, time.Local),
			time.Date(2017, 05, 29, 0, 0, 0, 0, time.Local)},
		// saturday after 15:00
		{time.Date(2017, 05, 27, 16, 0, 0, 0, time.Local),
			time.Date(2017, 05, 29, 16, 0, 0, 0, time.Local)},
		// sunday before 15:00
		{time.Date(2017, 06, 4, 0, 0, 0, 0, time.Local),
			time.Date(2017, 06, 5, 0, 0, 0, 0, time.Local)},
		// sunday after 15:00
		{time.Date(2017, 06, 4, 16, 0, 0, 0, time.Local),
			time.Date(2017, 06, 5, 16, 0, 0, 0, time.Local)},
	}
	for _, c := range cases {
		got := getDate(c.in)
		if got != c.expected {
			t.Errorf("getDate(%s) == %s, expected %s\n", c.in, got, c.expected)
		}
	}
}