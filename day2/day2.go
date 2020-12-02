package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var passwordLine = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

func Part1(passwordLines []string) (int, error) {
	var validCount int
	for _, l := range passwordLines {
		matches := passwordLine.FindStringSubmatch(l)
		if matches == nil {
			return 0, fmt.Errorf("day2: couldn't parse input line %q", l)
		}

		lower, err := strconv.Atoi(matches[1])
		if err != nil {
			return 0, fmt.Errorf("day2: parsing lower bound in %q: %w", l, err)
		}

		upper, err := strconv.Atoi(matches[2])
		if err != nil {
			return 0, fmt.Errorf("day2: parsing upper bound in %q: %w", l, err)
		}

		char, password := matches[3], matches[4]

		if count := strings.Count(password, char); lower <= count && count <= upper {
			validCount++
		}
	}
	return validCount, nil
}
