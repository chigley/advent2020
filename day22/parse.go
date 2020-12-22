package day22

import (
	"fmt"
	"strconv"
)

type Input struct {
	p1 []int
	p2 []int
}

func ParseInput(in [][]string) (*Input, error) {
	if len(in) != 2 {
		return nil, fmt.Errorf("day22: got %d input blocks, expected 2", len(in))
	}

	var (
		ret Input
		err error
	)

	ret.p1, err = parseBlock(in[0])
	if err != nil {
		return nil, fmt.Errorf("day22: parsing player 1: %w", err)
	}

	ret.p2, err = parseBlock(in[1])
	if err != nil {
		return nil, fmt.Errorf("day22: parsing player 2: %w", err)
	}

	return &ret, nil
}

func parseBlock(block []string) ([]int, error) {
	ret := make([]int, len(block)-1)
	for i := 1; i < len(block); i++ {
		var err error
		ret[i-1], err = strconv.Atoi(block[i])
		if err != nil {
			return nil, fmt.Errorf("day22: atoi: %w", err)
		}
	}
	return ret, nil
}
