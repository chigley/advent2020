package day19

func (in *Input) Part1() int {
	return in.matching()
}

func (in *Input) Part2() int {
	in.Rules[8] = OrRule{
		RefRule(42),
		AdjacentRule{RefRule(42), RefRule(8)},
	}

	in.Rules[11] = OrRule{
		AdjacentRule{RefRule(42), RefRule(31)},
		AdjacentRule{RefRule(42), RefRule(11), RefRule(31)},
	}

	return in.matching()
}

func (in *Input) matching() int {
	rule := in.Rules[0]
	cache := make(map[CacheKey]bool)

	var matching int
	for _, msg := range in.Messages {
		if rule.Match(in.Rules, cache, msg) {
			matching++
		}
	}
	return matching
}
