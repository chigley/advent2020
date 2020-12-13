package day13

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/deanveloper/modmath/v1/bigmod"
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

// I wouldn't have known to use CRT if not for /r/adventofcode and discussion
// with friends.
func Part2(busIDs []*int) (t *big.Int, err error) {
	// The bigmod library uses panics :(
	defer func() {
		if r := recover(); r != nil && err == nil {
			err = fmt.Errorf("day13: recovered: %v", r)
		}
	}()

	entries := make([]bigmod.CrtEntry, 0, len(busIDs))
	for i, busID := range busIDs {
		if busID == nil {
			continue
		}

		entries = append(entries, bigmod.CrtEntry{
			A: big.NewInt(int64(-i)),
			N: big.NewInt(int64(*busID)),
		})
	}
	return bigmod.SolveCrtMany(entries), nil
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
