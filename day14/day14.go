package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type Mask struct {
	andMask, orMask uint64
}

func (m *Mask) Apply(val uint64) uint64 {
	val &= m.andMask
	val |= m.orMask
	return val
}

func Part1(in []string) (uint64, error) {
	var mask *Mask
	memory := make(map[uint64]uint64)

	for _, l := range in {
		tokens := strings.SplitN(l, " = ", 2)
		if len(tokens) != 2 {
			return 0, fmt.Errorf("day14: can't parse line: %q", l)
		}

		lvalue, rvalue := tokens[0], tokens[1]
		if lvalue == "mask" {
			var err error
			mask, err = parseMask(rvalue)
			if err != nil {
				return 0, fmt.Errorf("day14: parsing mask %q: %w", rvalue, err)
			}
		} else {
			var addr uint64
			if _, err := fmt.Sscanf(lvalue, "mem[%d]", &addr); err != nil {
				return 0, fmt.Errorf("day14: can't parse token: %q: %w", l, err)
			}

			n, err := strconv.ParseUint(rvalue, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("day14: ParseUint: %w", err)
			}

			memory[addr] = mask.Apply(n)
		}
	}

	var sum uint64
	for _, n := range memory {
		sum += n
	}
	return sum, nil
}

func parseMask(mask string) (*Mask, error) {
	andMask, err := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 64)
	if err != nil {
		return nil, fmt.Errorf("day14: ParseUint: %w", err)
	}

	orMask, err := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 64)
	if err != nil {
		return nil, fmt.Errorf("day14: ParseUint: %w", err)
	}

	return &Mask{
		andMask: andMask,
		orMask:  orMask,
	}, nil
}
