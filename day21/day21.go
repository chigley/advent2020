package day21

type Allergens map[Allergen]Ingredients

func (as Allergens) ContainsNoAllergen(i Ingredient) bool {
	for _, is := range as {
		if _, ok := is[i]; ok {
			return false
		}
	}
	return true
}

type Ingredients map[Ingredient]struct{}

func (is Ingredients) Copy() Ingredients {
	ret := make(Ingredients, len(is))
	for i, val := range is {
		ret[i] = val
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

func Part1(foods []Food) int {
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

	var noAllergens int
	for _, f := range foods {
		for i := range f.ingredients {
			if allergens.ContainsNoAllergen(i) {
				noAllergens++
			}
		}
	}

	return noAllergens
}
