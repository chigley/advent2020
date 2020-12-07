package day7

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	regexpRule          = regexp.MustCompile(`^([a-z]+ [a-z]+) bags contain (.+)[.]$`)
	regexpContainedBags = regexp.MustCompile(`(\d+) ([a-z]+ [a-z]+) bags?`)
)

type (
	Colour string

	ColourSet   map[Colour]struct{}
	ContainedBy map[Colour]ColourSet
	Contains    map[Colour][]BagCount
)

type BagCount struct {
	colour Colour
	n      int
}

func Part1(containedBy ContainedBy, target Colour) int {
	var candidates []Colour
	for col := range containedBy[target] {
		candidates = append(candidates, col)
	}

	var n int
	seen := make(ColourSet)
	for len(candidates) > 0 {
		var col Colour
		col, candidates = candidates[0], candidates[1:]

		if _, ok := seen[col]; ok {
			continue
		}
		seen[col] = struct{}{}

		n++
		for col := range containedBy[col] {
			candidates = append(candidates, col)
		}
	}

	return n
}

func Part2(contains Contains, target Colour) int {
	var n int
	for _, bag := range contains[target] {
		// It would be nice to add memoisation to the Part2(contains, bag.colour)
		// call in case we already did it.
		n += bag.n + bag.n*Part2(contains, bag.colour)
	}
	return n
}

func ParseRules(in []string) (Contains, ContainedBy, error) {
	contains, containedBy := make(Contains), make(ContainedBy)

	for _, l := range in {
		matches := regexpRule.FindStringSubmatch(l)
		if len(matches) != 3 {
			return nil, nil, fmt.Errorf("day7: unable to parse rule: %q", l)
		}

		containingColour := Colour(matches[1])
		if matches[2] == "no other bags" {
			contains[containingColour] = nil
		} else {
			innerMatches := regexpContainedBags.FindAllStringSubmatch(matches[2], -1)
			if len(innerMatches) == 0 {
				return nil, nil, fmt.Errorf("day7: unable to parse rule: %q", l)
			}

			for _, ms := range innerMatches {
				if len(ms) != 3 {
					return nil, nil, fmt.Errorf("day7: unable to parse rule: %q", l)
				}

				n, err := strconv.Atoi(ms[1])
				if err != nil {
					return nil, nil, fmt.Errorf("day7: unable to parse rule: %q: atoi: %w", l, err)
				}

				containedColour := Colour(ms[2])

				contains[containingColour] = append(contains[containingColour], BagCount{
					colour: containedColour,
					n:      n,
				})

				if containedBy[containedColour] == nil {
					containedBy[containedColour] = make(ColourSet)
				}
				containedBy[containedColour][containingColour] = struct{}{}
			}
		}
	}

	return contains, containedBy, nil
}
