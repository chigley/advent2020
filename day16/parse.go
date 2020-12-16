package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/chigley/advent2020"
)

var regexpFieldLine = regexp.MustCompile(`^([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)$`)

func ParseInput(groups [][]string) (*Input, error) {
	if len(groups) != 3 {
		return nil, fmt.Errorf("day16: expected 3 input chunks, got %d", len(groups))
	}

	fields, err := parseFields(groups[0])
	if err != nil {
		return nil, fmt.Errorf("day16: parsing fields: %w", err)
	}

	yourTicket, err := parseTicket(groups[1][1])
	if err != nil {
		return nil, fmt.Errorf("day16: parsing your ticket: %w", err)
	}

	nearbyTickets := make([]Ticket, len(groups[2])-1)
	for i, l := range groups[2][1:] {
		t, err := parseTicket(l)
		if err != nil {
			return nil, fmt.Errorf("day16: parsing nearby ticket %q: %w", l, err)
		}
		nearbyTickets[i] = t
	}

	return &Input{
		Fields:        fields,
		YourTicket:    yourTicket,
		NearbyTickets: nearbyTickets,
	}, nil
}

func parseFields(lines []string) (map[FieldName]FieldRanges, error) {
	fields := make(map[FieldName]FieldRanges)
	for _, l := range lines {
		matches := regexpFieldLine.FindStringSubmatch(l)
		if len(matches) != 6 {
			return nil, fmt.Errorf("day16: can't parse %q", l)
		}

		min1, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("day16: atoi: %w", err)
		}

		max1, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("day16: atoi: %w", err)
		}

		min2, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, fmt.Errorf("day16: atoi: %w", err)
		}

		max2, err := strconv.Atoi(matches[5])
		if err != nil {
			return nil, fmt.Errorf("day16: atoi: %w", err)
		}

		fields[FieldName(matches[1])] = []FieldRange{
			{min1, max1},
			{min2, max2},
		}
	}
	return fields, nil
}

func parseTicket(ticket string) (Ticket, error) {
	ints, err := advent2020.ReadIntLines(strings.NewReader(ticket))
	if err != nil {
		return nil, fmt.Errorf("day16: ReadIntLine: %w", err)
	}
	if len(ints) != 1 {
		return nil, fmt.Errorf("day16: ticket too long: %q", ticket)
	}
	return ints[0], nil
}
