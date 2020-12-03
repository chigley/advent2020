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
