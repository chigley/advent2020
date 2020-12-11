package advent2020

type XY struct {
	X, Y int
}

func (p1 XY) Add(p2 XY) XY {
	return XY{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func (p XY) Adjacent() []XY {
	adjacent := make([]XY, 0, 8)
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}
			adjacent = append(adjacent, p.Add(XY{X: x, Y: y}))
		}
	}
	return adjacent
}
