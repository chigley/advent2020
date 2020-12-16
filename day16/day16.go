package day16

import (
	"errors"
	"strings"
)

type Input struct {
	Fields        map[FieldName]FieldRanges
	YourTicket    Ticket
	NearbyTickets []Ticket
}

type FieldName string

type FieldRanges []FieldRange

func (fs FieldRanges) IsValid(n int) bool {
	for _, f := range fs {
		if f.IsValid(n) {
			return true
		}
	}
	return false
}

type FieldRange struct {
	Min, Max int
}

func (f FieldRange) IsValid(n int) bool {
	return f.Min <= n && n <= f.Max
}

type Ticket []int

// Part1 has the side-effect of removing invalid tickets from i.NearbyTickets.
func (in *Input) Part1() int {
	var invalidSum int

	var nextIndex int
	for _, t := range in.NearbyTickets {
		isValid := true

	value:
		for _, val := range t {
			for _, f := range in.Fields {
				if f.IsValid(val) {
					continue value
				}
			}
			invalidSum += val
			isValid = false
		}

		if isValid {
			in.NearbyTickets[nextIndex] = t
			nextIndex++
		}
	}
	in.NearbyTickets = in.NearbyTickets[:nextIndex]

	return invalidSum
}

// Part2 has the side effect of removing elements from in.Fields.
func (in *Input) Part2() (int, error) {
	assignments := make(map[int]FieldName)

	// For each field we have left to assign, we count how many possible indices
	// it would be valid as. At least for my input, there was always a field
	// with exactly one valid index. We assign this one and repeat until all
	// indices have been assigned to a field.
fieldToAssign:
	for len(in.Fields) > 0 {
		for fieldName, f := range in.Fields {
			candidates := make([]int, 0, len(in.YourTicket))
			for fieldIndex := 0; fieldIndex < len(in.YourTicket); fieldIndex++ {
				if _, alreadyAssigned := assignments[fieldIndex]; !alreadyAssigned && in.allTicketsValid(f, fieldIndex) {
					candidates = append(candidates, fieldIndex)
				}
			}

			if len(candidates) == 1 {
				assignments[candidates[0]] = fieldName
				delete(in.Fields, fieldName)
				continue fieldToAssign
			}
		}

		// We don't use advent2020.ErrNoResult here because there almost
		// certainly _is_ a valid answer, just not findable with our approach.
		return 0, errors.New("day16: greedy approach doesn't work for this input")
	}

	product := 1
	for i, fieldName := range assignments {
		if strings.HasPrefix(string(fieldName), "departure") {
			product *= in.YourTicket[i]
		}
	}
	return product, nil
}

func (in *Input) allTicketsValid(field FieldRanges, i int) bool {
	for _, t := range in.NearbyTickets {
		if !field.IsValid(t[i]) {
			return false
		}
	}
	return true
}
