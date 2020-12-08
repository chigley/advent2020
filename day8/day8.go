package day8

import (
	"errors"
	"fmt"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/machine"
)

type outcome int

const (
	infiniteLoop outcome = iota
	halt
)

func Part1(in []string) (int, error) {
	program, err := machine.ParseProgram(in)
	if err != nil {
		return 0, fmt.Errorf("day8: parsing program: %w", err)
	}

	m := machine.New(program)
	outcome, err := runUntilOutcome(m)
	if err != nil {
		return 0, fmt.Errorf("day8: running machine: %w", err)
	}
	if outcome != infiniteLoop {
		return 0, errors.New("day8: expected infinite loop")
	}
	return m.Acc, nil
}

func Part2(in []string) (int, error) {
	program, err := machine.ParseProgram(in)
	if err != nil {
		return 0, fmt.Errorf("day8: parsing program: %w", err)
	}

	for i, instr := range program {
		if instr.Opcode != machine.OpJmp && instr.Opcode != machine.OpNop {
			continue
		}

		clonedProgram := machine.CloneProgram(program)
		if instr.Opcode == machine.OpJmp {
			clonedProgram[i].Opcode = machine.OpNop
		} else {
			// instr.Opcode == machine.OpNop
			clonedProgram[i].Opcode = machine.OpJmp
		}

		m := machine.New(clonedProgram)
		outcome, err := runUntilOutcome(m)
		if err != nil {
			return 0, fmt.Errorf("day8: running machine: %w", err)
		}
		if outcome == halt {
			return m.Acc, nil
		}
	}

	return 0, advent2020.ErrNoResult
}

func runUntilOutcome(m *machine.Machine) (outcome, error) {
	seen := make(map[int]struct{})
	for {
		if _, ok := seen[m.PC]; ok {
			return infiniteLoop, nil
		}
		seen[m.PC] = struct{}{}

		if err := m.Step(); errors.Is(err, machine.ErrHalt) {
			return halt, nil
		} else if err != nil {
			return 0, fmt.Errorf("day8: stepping machine: %w", err)
		}
	}
}
