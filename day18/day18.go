package day18

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Part1(in []string) (int, error) {
	return sumLines(in, true)
}

func Part2(in []string) (int, error) {
	return sumLines(in, false)
}

func sumLines(in []string, samePrecedence bool) (int, error) {
	var sum int
	for _, l := range in {
		n, err := Eval(l, samePrecedence)
		if err != nil {
			return 0, fmt.Errorf("day18: evaluating %q: %w", l, err)
		}
		sum += n
	}
	return sum, nil
}

func Eval(expr string, samePrecedence bool) (int, error) {
	postfix, err := infixToPostfix(expr, samePrecedence)
	if err != nil {
		return 0, fmt.Errorf("day18: converting infix to postfix: %w", err)
	}

	n, err := postfix.Eval()
	if err != nil {
		return 0, fmt.Errorf("day18: evaluating postfix: %w", err)
	}

	return n, nil
}

type PostfixExpr []interface{}

// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func infixToPostfix(expr string, samePrecedence bool) (PostfixExpr, error) {
	var (
		output    PostfixExpr
		operators []byte
	)

	var precedence map[byte]int
	if samePrecedence {
		precedence = map[byte]int{'+': 0, '*': 0}
	} else {
		precedence = map[byte]int{'+': 1, '*': 0}
	}

	// We can get away with removing whitespace and doing this character by
	// character since numbers in the puzzle input are always single-digit.
	expr = strings.ReplaceAll(expr, " ", "")
	for i := 0; i < len(expr); i++ {
		tok := expr[i]
		if tok == '(' {
			operators = append([]byte{tok}, operators...)
		} else if tok == ')' {
			for len(operators) > 0 && operators[0] != '(' {
				var op uint8
				op, operators = operators[0], operators[1:]
				output = append(output, op)
			}
			if len(operators) == 0 || operators[0] != '(' {
				return nil, errors.New("day18: mismatched parentheses")
			}
			operators = operators[1:]
		} else if tok == '+' || tok == '*' {
			for len(operators) > 0 && operators[0] != '(' && precedence[operators[0]] >= precedence[tok] {
				var op byte
				op, operators = operators[0], operators[1:]
				output = append(output, op)
			}
			operators = append([]byte{tok}, operators...)
		} else {
			n, err := strconv.Atoi(string(tok))
			if err != nil {
				return nil, fmt.Errorf("day18: atoi: %w", err)
			}
			output = append(output, n)
		}
	}

	for len(operators) > 0 {
		var op byte
		op, operators = operators[0], operators[1:]
		output = append(output, op)
	}

	return output, nil
}

func (e PostfixExpr) Eval() (int, error) {
	var stack []int

	for _, tok := range e {
		switch t := tok.(type) {
		case byte:
			if len(stack) < 2 {
				return 0, fmt.Errorf("day18: not enough numbers on stack for %v operation", t)
			}

			var arg1, arg2 int
			arg1, arg2, stack = stack[0], stack[1], stack[2:]

			var n int
			switch t {
			case '+':
				n = arg1 + arg2
			case '*':
				n = arg1 * arg2
			default:
				return 0, fmt.Errorf("day18: unknown operator %v", t)
			}
			stack = append([]int{n}, stack...)
		case int:
			stack = append([]int{t}, stack...)
		default:
			return 0, fmt.Errorf("day18: unexpected postfix token type %T", t)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("day18: resulting stack had %d elements, expected 1", len(stack))
	}
	return stack[0], nil
}
