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

type ParseMaskFunc func(mask string) ([]Mask, error)

func Part1(in []string) (uint64, error) {
	return sumMemoryValues(in, parseSingleMask, true)
}

func Part2(in []string) (uint64, error) {
	return sumMemoryValues(in, permuteFloatingMask, false)
}

func sumMemoryValues(in []string, parseMask ParseMaskFunc, maskValue bool) (uint64, error) {
	var masks []Mask
	memory := make(map[uint64]uint64)

	for _, l := range in {
		tokens := strings.SplitN(l, " = ", 2)
		if len(tokens) != 2 {
			return 0, fmt.Errorf("day14: can't parse line: %q", l)
		}

		lvalue, rvalue := tokens[0], tokens[1]
		if lvalue == "mask" {
			var err error
			masks, err = parseMask(rvalue)
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

			if maskValue {
				memory[addr] = masks[0].Apply(n)
			} else {
				for _, m := range masks {
					memory[m.Apply(addr)] = n
				}
			}
		}
	}

	var sum uint64
	for _, n := range memory {
		sum += n
	}
	return sum, nil
}

func parseSingleMask(mask string) ([]Mask, error) {
	andMask, err := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 64)
	if err != nil {
		return nil, fmt.Errorf("day14: ParseUint: %w", err)
	}

	orMask, err := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 64)
	if err != nil {
		return nil, fmt.Errorf("day14: ParseUint: %w", err)
	}

	return []Mask{{andMask: andMask, orMask: orMask}}, nil
}

func permuteFloatingMask(mask string) ([]Mask, error) {
	return permuteFloatingMaskHelper(mask, Mask{andMask: (1 << 36) - 1}), nil
}

func permuteFloatingMaskHelper(maskStr string, maskAcc Mask) []Mask {
	for i, c := range maskStr {
		bit := uint64(1 << (len(maskStr) - (i + 1)))
		switch c {
		case '1':
			maskAcc.orMask |= bit
		case 'X':
			// One branch where we set bit to 0
			ret := permuteFloatingMaskHelper(maskStr[i+1:], Mask{
				andMask: maskAcc.andMask &^ bit,
				orMask:  maskAcc.orMask,
			})

			// Another branch where we set bit to 1
			return append(ret, permuteFloatingMaskHelper(maskStr[i+1:], Mask{
				andMask: maskAcc.andMask,
				orMask:  maskAcc.orMask | bit,
			})...)
		}
	}
	return []Mask{maskAcc}
}
