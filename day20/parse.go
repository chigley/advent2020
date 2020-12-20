package day20

import (
	"fmt"
)

func ParseTiles(groups [][]string) (map[int]Tile, error) {
	tiles := make(map[int]Tile, len(groups))
	for _, g := range groups {
		var tileID int
		if _, err := fmt.Sscanf(g[0], "Tile %d:", &tileID); err != nil {
			return nil, fmt.Errorf("day20: scanning tile ID: %w", err)
		}

		tile := make(Tile, len(g)-1)
		for y := 0; y < len(g)-1; y++ {
			tile[y] = []byte(g[y+1])
		}

		tiles[tileID] = tile
	}
	return tiles, nil
}
