package day19

func (in *Input) Part1() int {
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
