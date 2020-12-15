package day15

type (
	Said map[int]Turn
	Turn int
)

func (s Said) Insert(n int, t Turn) (Turn, bool) {
	prev, ok := s[n]
	s[n] = t
	return prev, ok
}

func Part1(in []int) int {
	return nthNumber(in, 2020)
}

func nthNumber(startingNumbers []int, n Turn) int {
	if int(n) <= len(startingNumbers) {
		return startingNumbers[n-1]
	}

	said := make(Said)
	for i, x := range startingNumbers {
		said.Insert(x, Turn(i+1))
	}

	var (
		seenBefore   bool
		seenBeforeAt Turn
	)
	for turn := Turn(len(startingNumbers) + 1); ; turn++ {
		var x int
		if seenBefore {
			x = int((turn - 1) - seenBeforeAt)
		}

		if turn == n {
			return x
		}

		seenBeforeAt, seenBefore = said.Insert(x, turn)
	}
}
