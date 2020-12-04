package day4

import (
	"fmt"
	"strings"
)

type Passport map[string]string

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

type validationFunc func(p Passport) (bool, error)

func ParsePassports(in []string) ([]Passport, error) {
	var passports []Passport

	p := make(Passport)
	for i, l := range in {
		if l == "" {
			passports = append(passports, p)
			p = make(Passport)
			continue
		}

		for _, f := range strings.Split(l, " ") {
			splitField := strings.SplitN(f, ":", 2)
			if len(splitField) != 2 {
				return nil, fmt.Errorf("day4: failed to parse field: %q", f)
			}
			p[splitField[0]] = splitField[1]
		}

		if i == len(in)-1 {
			passports = append(passports, p)
		}
	}

	return passports, nil
}

func Part1(ps []Passport) (int, error) {
	return validPassports(ps, hasRequiredFields)
}

func validPassports(ps []Passport, validFunc validationFunc) (int, error) {
	var validCount int
	for _, p := range ps {
		valid, err := validFunc(p)
		if err != nil {
			return 0, fmt.Errorf("day4: validating passport: %w", err)
		}
		if valid {
			validCount++
		}
	}
	return validCount, nil
}

func hasRequiredFields(p Passport) (bool, error) {
	for _, f := range requiredFields {
		if _, ok := p[f]; !ok {
			return false, nil
		}
	}
	return true, nil
}
