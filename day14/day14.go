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

func Part2(in []string) (uint64, error) {
	var masks []Mask
	memory := make(map[uint64]uint64)

	for _, l := range in {
		tokens := strings.SplitN(l, " = ", 2)
		if len(tokens) != 2 {
			return 0, fmt.Errorf("day14: can't parse line: %q", l)
		}

		lvalue, rvalue := tokens[0], tokens[1]
		if lvalue == "mask" {
			masks = permuteMasks(rvalue)
		} else {
			var addr uint64
			if _, err := fmt.Sscanf(lvalue, "mem[%d]", &addr); err != nil {
				return 0, fmt.Errorf("day14: can't parse token: %q: %w", l, err)
			}

			n, err := strconv.ParseUint(rvalue, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("day14: ParseUint: %w", err)
			}

			for _, m := range masks {
				memory[m.Apply(addr)] = n
			}
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

func permuteMasks(mask string) []Mask {
	return permuteMasksHelper(mask, Mask{andMask: (1 << 36) - 1}, 0)
}

func permuteMasksHelper(mask string, startMask Mask, startIndex int) []Mask {
	for i := startIndex; i < len(mask); i++ {
		switch mask[i] {
		case '1':
			startMask.orMask |= 1 << (35 - i)
		case 'X':
			// One branch where we set bit 35-i to 0
			ret := permuteMasksHelper(mask, Mask{
				andMask: startMask.andMask & ^(1 << (35 - i)),
				orMask:  startMask.orMask,
			}, i+1)

			// Another branch where we set bit 35-i to 1
			return append(ret, permuteMasksHelper(mask, Mask{
				andMask: startMask.andMask,
				orMask:  startMask.orMask | (1 << (35 - i)),
			}, i+1)...)
		}
	}
	return []Mask{startMask}
}
