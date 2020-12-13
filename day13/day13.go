package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	Earliest int
	BusIDs   []*int
}

func Part1(in *Input) (int, error) {
	bestBusID, bestWaitTime := -1, 0
	for _, busID := range in.BusIDs {
		if busID == nil {
			// Skip the "x" inputs
			continue
		}

		if bestBusID == -1 {
			bestBusID = *busID
			bestWaitTime = waitTime(in.Earliest, *busID)
			continue
		}

		if wait := waitTime(in.Earliest, *busID); wait < bestWaitTime {
			bestBusID = *busID
			bestWaitTime = wait
		}
	}
	return bestBusID * bestWaitTime, nil
}

func waitTime(earliest, busID int) int {
	mod := earliest % busID
	if mod == 0 {
		return 0
	}
	return busID - mod
}

func ParseInput(in []string) (*Input, error) {
	if len(in) != 2 {
		return nil, fmt.Errorf("day13: expected 2 input lines, got %d", len(in))
	}

	earliest, err := strconv.Atoi(in[0])
	if err != nil {
		return nil, fmt.Errorf("day13: earliest atoi: %w", err)
	}

	busIDStrings := strings.Split(in[1], ",")
	busIDs := make([]*int, len(busIDStrings))
	for i, busIDStr := range busIDStrings {
		if busIDStr == "x" {
			busIDs[i] = nil
			continue
		}

		busID, err := strconv.Atoi(busIDStr)
		if err != nil {
			return nil, fmt.Errorf("day13: bus ID atoi: %w", err)
		}
		busIDs[i] = &busID
	}

	return &Input{
		Earliest: earliest,
		BusIDs:   busIDs,
	}, nil
}
