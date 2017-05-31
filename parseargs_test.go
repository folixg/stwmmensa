package main

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
		got := locationValid(c.in)
		if got != c.expected {
			t.Errorf("locationValid(%q) == %t, expected %t\n", c.in, got, c.expected)
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
		got := formatValid(c.in)
		if got != c.expected {
			t.Errorf("formatValid(%q) == %t, expected %t\n", c.in, got, c.expected)
		}
	}
}

func TestParseargs(t *testing.T) {
	cases := []struct {
		in []string
	}{
		{[]string{"stwmmensa", "-o", "test.xml", "-l", "411", "-f", "lis"}},
		{[]string{"stwmmensa", "--output=test.xml", "--location=411", "--format=lis"}},
	}
	for _, c := range cases {
		r := parseArgs(c.in)
		if !(r.format == "lis" && r.output == "test.xml" && r.location == "411") {
			t.Error("Failed to parse arguments")
		}
	}
	// default case
	r := parseArgs([]string{"stwmmensa", "-o", "test.xml"})
	if !(r.format == "xml" && r.output == "test.xml" && r.location == "421") {
		t.Error("Failed to parse arguments")
	}
}
