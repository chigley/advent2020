package day22

import (
	"errors"
	"strconv"
	"strings"
)

type Player int

const (
	Player1 Player = iota + 1
	Player2
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

func (in *Input) String() string {
	var b strings.Builder
	b.WriteString(cardsString(in.p1))
	b.WriteRune('\n')
	b.WriteString(cardsString(in.p2))
	return b.String()
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

func Part2(in *Input) (int, error) {
	winner, err := in.playRecursive()
	if err != nil {
		return 0, err
	}

	if winner == Player1 {
		return score(in.p1), nil
	}
	return score(in.p2), nil
}

func (in *Input) playRecursive() (Player, error) {
	seen := make(map[string]struct{})
	for len(in.p1) > 0 && len(in.p2) > 0 {
		key := in.String()
		if _, ok := seen[key]; ok {
			return Player1, nil
		}
		seen[key] = struct{}{}

		var p1Card, p2Card int
		p1Card, in.p1 = in.p1[0], in.p1[1:]
		p2Card, in.p2 = in.p2[0], in.p2[1:]

		if len(in.p1) < p1Card || len(in.p2) < p2Card {
			switch {
			case p1Card > p2Card:
				in.p1 = append(in.p1, p1Card, p2Card)
				continue
			case p2Card > p1Card:
				in.p2 = append(in.p2, p2Card, p1Card)
				continue
			default:
				return 0, errors.New("day22: tie?")
			}
		}

		nextIn := Input{
			p1: append([]int(nil), in.p1[:p1Card]...),
			p2: append([]int(nil), in.p2[:p2Card]...),
		}

		winner, err := nextIn.playRecursive()
		if err != nil {
			return 0, err
		}

		if winner == Player1 {
			in.p1 = append(in.p1, p1Card, p2Card)
		} else {
			in.p2 = append(in.p2, p2Card, p1Card)
		}
	}

	if len(in.p1) == 0 {
		return Player2, nil
	}
	return Player1, nil
}

func score(cards []int) int {
	var score int

	numCards := len(cards)
	for i, n := range cards {
		score += n * (numCards - i)
	}

	return score
}

func cardsString(cards []int) string {
	cardStrs := make([]string, len(cards))
	for i, c := range cards {
		cardStrs[i] = strconv.Itoa(c)
	}
	return strings.Join(cardStrs, ",")
}
