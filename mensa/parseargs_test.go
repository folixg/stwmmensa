package mensa

import "testing"

func TestLocation(t *testing.T) {
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

func TestFormat(t *testing.T) {
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

func TestParseargs(t *testing.T) {

	// good cases
	cases := []struct {
		in []string
	}{
		{[]string{"stwmmensa", "-o", "test.xml", "-l", "411", "-f", "lis"}},
		{[]string{"stwmmensa", "--output=test.xml", "--location=411", "--format=lis"}},
	}
	for _, c := range cases {
		r, err := ParseArgs(c.in)
		if !(r.Format == "lis" && r.Outfile == "test.xml" && r.Location == "411") {
			t.Error("Failed to parse arguments")
		}
		if err != nil {
			t.Errorf("%q Error was raised when parsing arguments", err)
		}
	}

	// check default values
	r, err := ParseArgs([]string{"stwmmensa", "-o", "test.xml"})
	if !(r.Format == "xml" && r.Outfile == "test.xml" && r.Location == "421") {
		t.Error("Failed to parse arguments")
	}
	if err != nil {
		t.Errorf("%q Error was raised when parsing arguments", err)
	}

	// missing output file
	_, err = ParseArgs([]string{"stwmmensa", "-f", "lis"})
	if err == nil {
		t.Error("Failed to detect missing output file.")
	}

	// illegal location
	_, err = ParseArgs([]string{"stwmmensa", "-o", "test", "-l", "487"})
	if err == nil {
		t.Error("Failed to detect illegal location ID.")
	}

	// illegal format
	_, err = ParseArgs([]string{"stwmmensa", "-o", "test", "-f", "html"})
	if err == nil {
		t.Error("Failed to detect illegal format.")
	}

	// print help
	_, err = ParseArgs([]string{"stwmmensa", "-h"})
	if err == nil {
		t.Error("Failed to show help.")
	}

	// unknown option
	_, err = ParseArgs([]string{"stwmmensa", "-o", "test", "--foo=bar"})
	if err == nil {
		t.Error("Failed to detect unknown option.")
	}
}
