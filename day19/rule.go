package day19

import (
	"fmt"
	"strconv"
	"strings"
)

type CacheKey struct {
	// It is poor that we keep regenerating this with String() calls. It might
	// be nicer for each rule to become a struct, which stores its original
	// string representation alongside the rule.
	RuleDef string

	Message string
}

type RuleID int

type Rule interface {
	Match(rules map[RuleID]Rule, cache map[CacheKey]bool, message string) bool
	fmt.Stringer
}

type AdjacentRule []Rule

func (r AdjacentRule) Match(rules map[RuleID]Rule, cache map[CacheKey]bool, message string) (ret bool) {
	key := CacheKey{
		RuleDef: r.String(),
		Message: message,
	}
	if ret, ok := cache[key]; ok {
		return ret
	}
	defer func() {
		cache[key] = ret
	}()

	if len(r) == 1 {
		return r[0].Match(rules, cache, message)
	}

	for i := 1; i < len(message); i++ {
		if r[0].Match(rules, cache, message[:i]) && r[1:].Match(rules, cache, message[i:]) {
			return true
		}
	}
	return false
}

func (r AdjacentRule) String() string {
	subrules := make([]string, len(r))
	for i, subrule := range r {
		subrules[i] = subrule.String()
	}
	return strings.Join(subrules, " ")
}

type OrRule []Rule

func (r OrRule) Match(rules map[RuleID]Rule, cache map[CacheKey]bool, message string) (ret bool) {
	key := CacheKey{
		RuleDef: r.String(),
		Message: message,
	}
	if ret, ok := cache[key]; ok {
		return ret
	}
	defer func() {
		cache[key] = ret
	}()

	for _, subrule := range r {
		if subrule.Match(rules, cache, message) {
			return true
		}
	}
	return false
}

func (r OrRule) String() string {
	subrules := make([]string, len(r))
	for i, subrule := range r {
		subrules[i] = subrule.String()
	}
	return strings.Join(subrules, " | ")
}

type TerminalRule byte

func (r TerminalRule) Match(_ map[RuleID]Rule, _ map[CacheKey]bool, message string) bool {
	// Don't bother with a cache lookup here: it probably isn't worth the String() call
	return len(message) == 1 && message[0] == byte(r)
}

func (r TerminalRule) String() string {
	return fmt.Sprintf(`"%c"`, r)
}

type RefRule RuleID

func (r RefRule) Match(rules map[RuleID]Rule, cache map[CacheKey]bool, message string) (ret bool) {
	key := CacheKey{
		RuleDef: r.String(),
		Message: message,
	}
	if ret, ok := cache[key]; ok {
		return ret
	}
	defer func() {
		cache[key] = ret
	}()

	return rules[RuleID(r)].Match(rules, cache, message)
}

func (r RefRule) String() string {
	return strconv.Itoa(int(r))
}
