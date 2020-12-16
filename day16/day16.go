package day16

type Input struct {
	Fields        map[FieldName]FieldRanges
	YourTicket    Ticket
	NearbyTickets []Ticket
}

type FieldName string

type FieldRanges []FieldRange

func (fs FieldRanges) IsValid(n int) bool {
	for _, f := range fs {
		if f.IsValid(n) {
			return true
		}
	}
	return false
}

type FieldRange struct {
	Min, Max int
}

func (f FieldRange) IsValid(n int) bool {
	return f.Min <= n && n <= f.Max
}

type Ticket []int

// Part1 has the side-effect of removing invalid tickets from i.NearbyTickets.
func (in *Input) Part1() int {
	var invalidSum int

	var nextIndex int
	for _, t := range in.NearbyTickets {
		isValid := true

	value:
		for _, val := range t {
			for _, f := range in.Fields {
				if f.IsValid(val) {
					continue value
				}
			}
			invalidSum += val
			isValid = false
		}

		if isValid {
			in.NearbyTickets[nextIndex] = t
			nextIndex++
		}
	}
	in.NearbyTickets = in.NearbyTickets[:nextIndex]

	return invalidSum
}
