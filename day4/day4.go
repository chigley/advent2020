package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	regexpHeight     = regexp.MustCompile(`^(\d+)(cm|in)$`)
	regexpHairColour = regexp.MustCompile(`^#[a-f\d]{6}$`)
	regexpEyeColour  = regexp.MustCompile(`^(?:amb|blu|brn|gry|grn|hzl|oth)$`)
	regexpPassportID = regexp.MustCompile(`^\d{9}$`)
)

type Passport map[string]string

var requiredFields = map[string]fieldValidationFunc{
	"byr": validFourDigitRange(1920, 2002),
	"iyr": validFourDigitRange(2010, 2020),
	"eyr": validFourDigitRange(2020, 2030),
	"hgt": validHeight,
	"hcl": func(val string) (bool, error) { return regexpHairColour.MatchString(val), nil },
	"ecl": func(val string) (bool, error) { return regexpEyeColour.MatchString(val), nil },
	"pid": func(val string) (bool, error) { return regexpPassportID.MatchString(val), nil },
}

type (
	passwordValidationFunc func(p Passport) (bool, error)
	fieldValidationFunc    func(val string) (bool, error)
)

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

func Part2(ps []Passport) (int, error) {
	return validPassports(ps, hasValidRequiredFields)
}

func validPassports(ps []Passport, validFunc passwordValidationFunc) (int, error) {
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
	for f := range requiredFields {
		if _, ok := p[f]; !ok {
			return false, nil
		}
	}
	return true, nil
}

func hasValidRequiredFields(p Passport) (bool, error) {
	for name, validFunc := range requiredFields {
		value, ok := p[name]
		if !ok {
			return false, nil
		}

		valid, err := validFunc(value)
		if err != nil {
			return false, fmt.Errorf("day4: validating %q field: %w", name, err)
		}
		if !valid {
			return false, nil
		}
	}
	return true, nil
}

func validFourDigitRange(lower, upper int) fieldValidationFunc {
	return func(val string) (bool, error) {
		if len(val) != 4 {
			return false, nil
		}

		valInt, err := strconv.Atoi(val)
		if err != nil {
			return false, fmt.Errorf("day4: atoi: %w", err)
		}

		valid := lower <= valInt && valInt <= upper
		return valid, nil
	}
}

func validHeight(val string) (bool, error) {
	matches := regexpHeight.FindStringSubmatch(val)
	if len(matches) != 3 {
		return false, nil
	}

	height, err := strconv.Atoi(matches[1])
	if err != nil {
		return false, fmt.Errorf("day4: atoi: %w", err)
	}

	unit := matches[2]
	switch unit {
	case "cm":
		return 150 <= height && height <= 193, nil
	case "in":
		return 59 <= height && height <= 76, nil
	default:
		return false, fmt.Errorf("day4: don't know how to validate height unit %q", unit)
	}
}
