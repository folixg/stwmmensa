package mensa

import "testing"

func TestDishes(t *testing.T) {
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
