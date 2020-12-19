package day19

import (
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	Rules    map[RuleID]Rule
	Messages []string
}

func ParseInput(in [][]string) (*Input, error) {
	if len(in) != 2 {
		return nil, fmt.Errorf("day19: expected 2 input blocks, got %d", len(in))
	}

	rules := make(map[RuleID]Rule)
	for _, r := range in[0] {
		toks := strings.SplitN(r, ": ", 2)
		if len(toks) != 2 {
			return nil, fmt.Errorf("day19: can't parse rule: %q", r)
		}

		ruleID, err := strconv.Atoi(toks[0])
		if err != nil {
			return nil, fmt.Errorf("day19: atoi: %w", err)
		}

		rules[RuleID(ruleID)], err = parseRule(toks[1])
		if err != nil {
			return nil, fmt.Errorf("day19: parsing rule %q: %w", toks[1], err)
		}
	}

	return &Input{
		Rules:    rules,
		Messages: in[1],
	}, nil
}

func parseRule(rule string) (Rule, error) {
	if strings.Contains(rule, `"`) {
		return TerminalRule(rule[1]), nil
	}

	if strings.Contains(rule, "|") {
		toks := strings.Split(rule, " | ")
		subrules := make(OrRule, len(toks))
		for i := 0; i < len(toks); i++ {
			var err error
			subrules[i], err = parseRule(toks[i])
			if err != nil {
				return nil, err
			}
		}
		return subrules, nil
	}

	if strings.Contains(rule, " ") {
		toks := strings.Split(rule, " ")
		subrules := make(AdjacentRule, len(toks))
		for i := 0; i < len(toks); i++ {
			var err error
			subrules[i], err = parseRule(toks[i])
			if err != nil {
				return nil, err
			}
		}
		return subrules, nil
	}

	ruleID, err := strconv.Atoi(rule)
	if err != nil {
		return nil, fmt.Errorf("day19: atoi: %w", err)
	}
	return RefRule(ruleID), nil
}
