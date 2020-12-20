package day20

import (
	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/bfs"
)

type BFSNode struct {
	picture      Picture
	tilesToPlace map[int]Tile
	noncer       *Noncer
}

func (b *BFSNode) IsGoal() bool {
	return b.picture.NextEmptySquare() == nil
}

func (b *BFSNode) Neighbours() ([]bfs.Node, error) {
	var neighbours []bfs.Node

	squareToFill := *b.picture.NextEmptySquare()
	for tileID, baseTile := range b.tilesToPlace {
		for _, tile := range baseTile.Permutations() {
			if b.picture.Fits(tile, squareToFill) {
				neighbours = append(neighbours, b.placeTile(tile, tileID, squareToFill))
			}
		}
	}

	return neighbours, nil
}

// We don't need to worry about hitting the same state twice. Stop the BFS
// implementation from worrying about this with this lovely little bodge.
func (b *BFSNode) Key() interface{} {
	return b.noncer.Nonce()
}

func (b *BFSNode) placeTile(tile Tile, tileID int, pos advent2020.XY) *BFSNode {
	tilesToPlace := make(map[int]Tile)
	for tileID, t := range b.tilesToPlace {
		tilesToPlace[tileID] = t
	}
	delete(tilesToPlace, tileID)

	return &BFSNode{
		picture:      b.picture.Place(tile, tileID, pos),
		tilesToPlace: tilesToPlace,
		noncer:       b.noncer,
	}
}
