package day22

import (
	"errors"
)

func (in *Input) Copy() *Input {
	p1Copy := make([]int, len(in.p1))
	copy(p1Copy, in.p1)

	p2Copy := make([]int, len(in.p2))
	copy(p2Copy, in.p2)

	return &Input{
		p1: p1Copy,
		p2: p2Copy,
	}
}

func Part1(in *Input) (int, error) {
	for len(in.p1) > 0 && len(in.p2) > 0 {
		var p1Card, p2Card int
		p1Card, in.p1 = in.p1[0], in.p1[1:]
		p2Card, in.p2 = in.p2[0], in.p2[1:]

		switch {
		case p1Card > p2Card:
			in.p1 = append(in.p1, p1Card, p2Card)
		case p2Card > p1Card:
			in.p2 = append(in.p2, p2Card, p1Card)
		default:
			return 0, errors.New("day22: tie?")
		}
	}

	if len(in.p1) == 0 {
		return score(in.p2), nil
	}
	return score(in.p1), nil
}

func score(cards []int) int {
	var score int

	numCards := len(cards)
	for i, n := range cards {
		score += n * (numCards - i)
	}

	return score
}
