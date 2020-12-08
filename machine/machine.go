package machine

import (
	"errors"
	"fmt"
)

var ErrHalt = errors.New("machine: halt")

type Machine struct {
	program []Instruction
	PC      int
	Acc     int
}

func New(program []Instruction) *Machine {
	return &Machine{
		program: program,
		PC:      0,
		Acc:     0,
	}
}

func (m *Machine) Step() error {
	if m.PC == len(m.program) {
		return ErrHalt
	}

	instr := m.program[m.PC]
	m.PC++

	switch instr.Opcode {
	case OpAcc:
		m.Acc += instr.Arg
	case OpJmp:
		m.PC += instr.Arg - 1
	case OpNop:
	default:
		return fmt.Errorf("machine: unknown opcode %d", instr.Opcode)
	}

	return nil
}
