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

type HexDirection int

const (
	HexEast HexDirection = iota
	HexSouthEast
	HexSouthWest
	HexWest
	HexNorthWest
	HexNorthEast
)

// https://www.redblobgames.com/grids/hexagons/#coordinates-cube
var HexCompass = map[HexDirection]XYZ{
	HexEast:      {1, -1, 0},
	HexSouthEast: {0, -1, 1},
	HexSouthWest: {-1, 0, 1},
	HexWest:      {-1, 1, 0},
	HexNorthWest: {0, 1, -1},
	HexNorthEast: {1, 0, -1},
}

type XYZ struct {
	X, Y, Z int
}

func (p1 XYZ) Add(p2 XYZ) XYZ {
	return XYZ{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
		Z: p1.Z + p2.Z,
	}
}

func (p XYZ) AddHexCompass(d HexDirection) XYZ {
	return p.Add(HexCompass[d])
}

func (p XYZ) HexAdjacent() []XYZ {
	adjacent := make([]XYZ, len(HexCompass))
	var i int
	for _, d := range HexCompass {
		adjacent[i] = p.Add(d)
		i++
	}
	return adjacent
}

type WXYZ struct {
	W, X, Y, Z int
}

func (p1 WXYZ) Add(p2 WXYZ) WXYZ {
	return WXYZ{
		W: p1.W + p2.W,
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
		Z: p1.Z + p2.Z,
	}
}

func (p WXYZ) Adjacent() []WXYZ {
	adjacent := make([]WXYZ, 0, 80)
	for w := -1; w <= 1; w++ {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					if w == 0 && x == 0 && y == 0 && z == 0 {
						continue
					}
					adjacent = append(adjacent, p.Add(WXYZ{w, x, y, z}))
				}
			}
		}
	}
	return adjacent
}
