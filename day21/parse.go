package day21

import (
	"fmt"
	"regexp"
	"strings"
)

var regexpFoodLine = regexp.MustCompile(`^(.+) \(contains (.+)\)$`)

type (
	Ingredient string
	Allergen   string
)

type Food struct {
	ingredients       Ingredients
	definiteAllergens []Allergen
}

func ParseInput(in []string) ([]Food, error) {
	var foods []Food
	for _, l := range in {
		matches := regexpFoodLine.FindStringSubmatch(l)
		if len(matches) != 3 {
			return nil, fmt.Errorf("day21: can't parse line: %q", l)
		}

		ingredients := make(Ingredients)
		for _, i := range strings.Split(matches[1], " ") {
			ingredients[Ingredient(i)] = struct{}{}
		}

		allergenStrs := strings.Split(matches[2], ", ")
		allergens := make([]Allergen, len(allergenStrs))
		for i, a := range allergenStrs {
			allergens[i] = Allergen(a)
		}

		foods = append(foods, Food{
			ingredients:       ingredients,
			definiteAllergens: allergens,
		})
	}
	return foods, nil
}
