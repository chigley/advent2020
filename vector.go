package advent2020

var Directions = []XY{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

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
	adjacent := make([]XY, len(Directions))
	for i, d := range Directions {
		adjacent[i] = p.Add(d)
	}
	return adjacent
}
