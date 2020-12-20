package day20

func Part1(tiles map[int]Tile) int {
	edgeCount := make(map[string]int)
	for _, tile := range tiles {
		for _, e := range tile.Edges() {
			edgeCount[e]++
		}
	}

	var cornerIDs []int
	for tileID, tile := range tiles {
		var numUniqueEdges int
		for _, e := range tile.Edges() {
			if edgeCount[e] == 1 {
				numUniqueEdges++
			}
		}

		if numUniqueEdges == 2 {
			cornerIDs = append(cornerIDs, tileID)
		}
	}

	product := 1
	for _, id := range cornerIDs {
		product *= id
	}
	return product
}
