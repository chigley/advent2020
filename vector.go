package advent2020

type CompassDirection int

const (
	North CompassDirection = iota
	East
	South
	West
)

var Compass = map[CompassDirection]XY{
	North: {0, 1},
	East:  {1, 0},
	South: {0, -1},
	West:  {-1, 0},
}

var Directions = []XY{
	{-1, -1},
	Compass[West],
	{-1, 1},
	Compass[South],
	Compass[North],
	{1, -1},
	Compass[East],
	{1, 1},
}

func (d CompassDirection) RotateClockwise(steps int) CompassDirection {
	return (d + CompassDirection(steps)) % 4
}

func (d CompassDirection) RotateAnticlockwise(steps int) CompassDirection {
	return d.RotateClockwise(4 - steps%4)
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

func (p XY) Multiply(n int) XY {
	return XY{
		X: n * p.X,
		Y: n * p.Y,
	}
}

func (p XY) Manhattan() int {
	return Abs(p.X) + Abs(p.Y)
}

func (p XY) Adjacent() []XY {
	adjacent := make([]XY, len(Directions))
	for i, d := range Directions {
		adjacent[i] = p.Add(d)
	}
	return adjacent
}

func (p XY) RotateClockwise(steps int) XY {
	for i := 0; i < steps%4; i++ {
		tmp := p.Y
		p.Y = -p.X
		p.X = tmp
	}
	return p
}

func (p XY) RotateAnticlockwise(steps int) XY {
	return p.RotateClockwise(4 - steps%4)
}
