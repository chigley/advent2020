package day20

import (
	"fmt"
	"math"
	"strings"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/bfs"
)

func BuildPicture(tiles map[int]Tile) (*Picture, error) {
	var tileSize int
	for _, tile := range tiles {
		tileSize = len(tile)
		break
	}

	size := int(math.Sqrt(float64(len(tiles))))
	start := &BFSNode{
		picture:      NewPicture(size, tileSize),
		tilesToPlace: tiles,
		noncer:       NewNoncer(),
	}

	path, err := bfs.Search(start)
	if err != nil {
		return nil, fmt.Errorf("day20: BFS: %w", err)
	}

	return &path[len(path)-1].(*BFSNode).picture, nil
}

func Part1(p *Picture) int {
	return p.tileIDs[0][0] * p.tileIDs[0][p.size-1] * p.tileIDs[p.size-1][0] * p.tileIDs[p.size-1][p.size-1]
}

func Part2(p *Picture) (int, error) {
	image := strings.Split(p.ImageString(false), "\n")
	tile := make(Tile, len(image))
	for i := 0; i < len(image); i++ {
		tile[i] = []byte(image[i])
	}

	for _, t := range tile.Permutations() {
		if ok := t.ReplaceSeaMonsters(); !ok {
			continue
		}

		var count int
		for y := 0; y < len(t); y++ {
			for x := 0; x < len(t[0]); x++ {
				if t[y][x] == '#' {
					count++
				}
			}
		}
		return count, nil
	}

	return 0, advent2020.ErrNoResult
}
