package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var passwordLine = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

type validationFunc func(char uint8, tok1, tok2 int, password string) bool

func Part1(passwordLines []string) (int, error) {
	return validPasswords(passwordLines, validLetterCount)
}

func Part2(passwordLines []string) (int, error) {
	return validPasswords(passwordLines, oneOccurrenceInPositions)
}

func validPasswords(passwordLines []string, validFunc validationFunc) (int, error) {
	var validCount int
	for _, l := range passwordLines {
		matches := passwordLine.FindStringSubmatch(l)
		if matches == nil {
			return 0, fmt.Errorf("day2: couldn't parse input line %q", l)
		}

		tok1, err := strconv.Atoi(matches[1])
		if err != nil {
			return 0, fmt.Errorf("day2: parsing first token in %q: %w", l, err)
		}

		tok2, err := strconv.Atoi(matches[2])
		if err != nil {
			return 0, fmt.Errorf("day2: parsing second token in %q: %w", l, err)
		}

		char, password := matches[3][0], matches[4]
		if validFunc(char, tok1, tok2, password) {
			validCount++
		}
	}
	return validCount, nil
}

func validLetterCount(char uint8, lower, upper int, password string) bool {
	count := strings.Count(password, string(char))
	return lower <= count && count <= upper
}

func oneOccurrenceInPositions(char uint8, i1, i2 int, password string) bool {
	i1--
	i2--
	return (password[i1] == char && password[i2] != char) || (password[i1] != char && password[i2] == char)
}
