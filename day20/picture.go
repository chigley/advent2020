package day20

import (
	"strings"

	"github.com/chigley/advent2020"
)

type Picture struct {
	// We keep state on what the tiles look like _and_ their IDs. Tile IDs alone
	// aren't sufficient because they can't tell us which translations were
	// performed.
	tiles   [][]Tile
	tileIDs [][]int

	// Width/height measured in number of tiles
	size int

	// Width/height of each individual tile
	tileSize int
}

func NewPicture(size, tileSize int) Picture {
	tiles := make([][]Tile, size)
	for i := 0; i < size; i++ {
		tiles[i] = make([]Tile, size)
	}

	tileIDs := make([][]int, size)
	for i := 0; i < size; i++ {
		tileIDs[i] = make([]int, size)
	}

	return Picture{
		tiles:    tiles,
		tileIDs:  tileIDs,
		size:     size,
		tileSize: tileSize,
	}
}

// We fill in the picture from left to right, top to bottom. Moving in this
// direction, NextEmptySquare returns the position of the next square to be
// filled, or nil if it is already full.
func (p Picture) NextEmptySquare() *advent2020.XY {
	for y := 0; y < p.size; y++ {
		for x := 0; x < p.size; x++ {
			if p.tiles[y][x] == nil {
				return &advent2020.XY{X: x, Y: y}
			}
		}
	}
	return nil
}

func (p Picture) Fits(t Tile, pos advent2020.XY) bool {
	leftNeighbour := p.TileAt(pos.Add(advent2020.XY{X: -1, Y: 0}))
	topNeighbour := p.TileAt(pos.Add(advent2020.XY{X: 0, Y: -1}))
	return (leftNeighbour == nil || leftNeighbour.RightSideMatchesLeftOf(t)) &&
		(topNeighbour == nil || topNeighbour.BottomSideMatchesTopOf(t))
}

func (p Picture) Clone() Picture {
	tiles := make([][]Tile, p.size)
	for y := 0; y < p.size; y++ {
		tiles[y] = make([]Tile, p.size)
		copy(tiles[y], p.tiles[y])
	}

	tileIDs := make([][]int, p.size)
	for y := 0; y < p.size; y++ {
		tileIDs[y] = make([]int, p.size)
		copy(tileIDs[y], p.tileIDs[y])
	}

	return Picture{
		tiles:    tiles,
		tileIDs:  tileIDs,
		size:     p.size,
		tileSize: p.tileSize,
	}
}

func (p Picture) Place(t Tile, tileID int, pos advent2020.XY) Picture {
	ret := p.Clone()
	ret.tiles[pos.Y][pos.X] = t
	ret.tileIDs[pos.Y][pos.X] = tileID
	return ret
}

func (p Picture) InBounds(pos advent2020.XY) bool {
	return 0 <= pos.X && pos.X < p.size && 0 <= pos.Y && pos.Y < p.size
}

func (p Picture) TileAt(pos advent2020.XY) Tile {
	if !p.InBounds(pos) {
		return nil
	}
	return p.tiles[pos.Y][pos.X]
}

func (p Picture) String() string {
	return p.ImageString(true)
}

func (p Picture) ImageString(includeBorders bool) string {
	var (
		minY, maxY int
		size       int
	)
	if includeBorders {
		minY, maxY = 0, p.tileSize
		size = p.tileSize
	} else {
		minY, maxY = 1, p.tileSize-1
		size = p.tileSize - 2
	}

	var b strings.Builder
	for tileY := 0; tileY < p.size; tileY++ {
		for y := minY; y < maxY; y++ {
			for tileX := 0; tileX < p.size; tileX++ {
				t := p.tiles[tileY][tileX]
				if t == nil {
					b.WriteString(strings.Repeat("?", size))
				} else {
					if includeBorders {
						b.WriteString(string(t[y]))
					} else {
						row := string(t[y])
						b.WriteString(row[1 : len(row)-1])
					}
				}

				if includeBorders && tileX != p.size-1 {
					b.WriteRune(' ')
				}
			}

			if !(tileY == p.size-1 && y == maxY-1) {
				b.WriteRune('\n')
			}
		}

		if includeBorders && tileY != p.size-1 {
			b.WriteRune('\n')
		}
	}
	return b.String()
}
