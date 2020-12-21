package day21

import (
	"sort"
	"strings"

	"github.com/chigley/advent2020"
)

type Allergens map[Allergen]Ingredients

func (as Allergens) ContainsNoAllergen(i Ingredient) bool {
	for _, is := range as {
		if _, ok := is[i]; ok {
			return false
		}
	}
	return true
}

func (as Allergens) DeleteIngredient(i Ingredient) {
	for _, is := range as {
		delete(is, i)
	}
}

type Ingredients map[Ingredient]struct{}

func (is Ingredients) Copy() Ingredients {
	ret := make(Ingredients, len(is))
	for i, val := range is {
		ret[i] = val
	}
	return ret
}

func (is Ingredients) Slice() []Ingredient {
	ret := make([]Ingredient, len(is))
	var count int
	for i := range is {
		ret[count] = i
		count++
	}
	return ret
}

func (is1 Ingredients) Intersect(is2 Ingredients) {
	for i := range is1 {
		if _, ok := is2[i]; !ok {
			delete(is1, i)
		}
	}
}

func Day21(foods []Food) (int, string, error) {
	allergens := make(Allergens)

	for _, f := range foods {
		for _, a := range f.definiteAllergens {
			if ingredientsContainingA, ok := allergens[a]; ok {
				ingredientsContainingA.Intersect(f.ingredients)
			} else {
				allergens[a] = f.ingredients.Copy()
			}
		}
	}

	// Part 1
	var (
		ingredientsWithAllergens = make(Ingredients)
		noAllergens              int
	)
	for _, f := range foods {
		for i := range f.ingredients {
			if allergens.ContainsNoAllergen(i) {
				noAllergens++
			} else {
				ingredientsWithAllergens[i] = struct{}{}
			}
		}
	}

	// Part 2
	ingredientToAllergen := make(map[Ingredient]Allergen)
outer:
	for len(ingredientsWithAllergens) > 0 {
		for a, is := range allergens {
			if len(is) == 1 {
				i := is.Slice()[0]

				ingredientToAllergen[i] = a

				delete(ingredientsWithAllergens, i)
				allergens.DeleteIngredient(i)

				continue outer
			}
		}
		return 0, "", advent2020.ErrNoResult
	}

	ingredients := make([]string, len(ingredientToAllergen))
	var count int
	for i := range ingredientToAllergen {
		ingredients[count] = string(i)
		count++
	}

	sort.Slice(ingredients, func(i, j int) bool {
		return ingredientToAllergen[Ingredient(ingredients[i])] < ingredientToAllergen[Ingredient(ingredients[j])]
	})
	return noAllergens, strings.Join(ingredients, ","), nil
}
