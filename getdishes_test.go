package main

import "testing"

func TestDishes(t *testing.T) {
	expected := []dish{
		{category: "Tagesgericht 1",
			name: "Zartweizen mit Tomaten, Zucchini und Auberginen"},
		{category: "Tagesgericht 3",
			name: "Cevapcici von der Pute mit Ajvar"},
		{category: "Aktionsessen 4",
			name: "Münchner Biergulasch (GQB) (R)(99)"},
		{category: "Self-Service",
			name: "Zartweizen mit Tomaten, Zucchini und Auberginen"},
		{category: "Self-Service",
			name: "Rigatoni mit Paprikapesto"},
		{category: "Self-Service",
			name: "Bunte Nudel-Hackfleisch-Pfanne (R)"},
	}
	got := getDishes("http://www.studentenwerk-muenchen.de/mensa/speiseplan/speiseplan_2017-05-22_421_-de.html")
	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("Expected: %s, %s\n Got: %s, %s", expected[i].category,
				expected[i].name, got[i].category, got[i].name)
		}
	}
}
