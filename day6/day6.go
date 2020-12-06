package day6

type (
	Question rune

	PersonIndex int
	People      map[PersonIndex]struct{}
)

type Group struct {
	answeredBy map[Question]People
	people     int
}

func NewGroup() Group {
	return Group{
		answeredBy: make(map[Question]People),
		people:     0,
	}
}

func (g Group) AnsweredByAll() int {
	var questions int
	for _, answerers := range g.answeredBy {
		if len(answerers) == g.people {
			questions++
		}
	}
	return questions
}

func ParseGroups(in [][]string) []Group {
	groups := make([]Group, len(in))
	for i, lines := range in {
		g := NewGroup()
		for _, l := range lines {
			for _, c := range l {
				q := Question(c)
				if g.answeredBy[q] == nil {
					g.answeredBy[q] = make(People)
				}
				g.answeredBy[q][PersonIndex(g.people)] = struct{}{}
			}
			g.people++
		}
		groups[i] = g
	}
	return groups
}

func Part1(groups []Group) int {
	var sum int
	for _, g := range groups {
		sum += len(g.answeredBy)
	}
	return sum
}

func Part2(groups []Group) int {
	var sum int
	for _, g := range groups {
		sum += g.AnsweredByAll()
	}
	return sum
}
