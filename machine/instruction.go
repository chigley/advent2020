package machine

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	Opcode Opcode
	Arg    int
}

func ParseProgram(program []string) ([]Instruction, error) {
	instructions := make([]Instruction, len(program))
	for i, instruction := range program {
		parsedInstruction, err := parseInstruction(instruction)
		if err != nil {
			return nil, fmt.Errorf("machine: parsing instruction %q: %w", instruction, err)
		}
		instructions[i] = *parsedInstruction
	}
	return instructions, nil
}

func parseInstruction(in string) (*Instruction, error) {
	tokens := strings.SplitN(in, " ", 2)
	if len(tokens) != 2 {
		return nil, fmt.Errorf("machine: got %d tokens, expected 2", len(tokens))
	}

	opcode, ok := ops[tokens[0]]
	if !ok {
		return nil, fmt.Errorf("machine: unknown opcode: %q", tokens[0])
	}

	arg, err := strconv.Atoi(tokens[1])
	if err != nil {
		return nil, fmt.Errorf("machine: parsing argument %q: %w", tokens[1], err)
	}

	return &Instruction{
		Opcode: opcode,
		Arg:    arg,
	}, nil
}

func CloneProgram(program []Instruction) []Instruction {
	ret := make([]Instruction, len(program))
	copy(ret, program)
	return ret
}
