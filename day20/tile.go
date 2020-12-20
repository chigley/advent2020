package day20

var seaMonster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

const (
	seaMonsterHeight = 3
	seaMonsterWidth  = len("#    ##    ##    ###")
)

type Tile [][]byte

func (t Tile) Permutations() []Tile {
	return append(t.RotatePermutations(), t.Flip().RotatePermutations()...)
}

func (t Tile) RotatePermutations() []Tile {
	ret := make([]Tile, 4)
	ret[0] = t
	for i := 0; i < 3; i++ {
		t = t.RotateLeft()
		ret[i+1] = t
	}
	return ret
}

func (t Tile) RotateLeft() Tile {
	ret := make(Tile, len(t))
	for i := 0; i < len(t); i++ {
		ret[i] = make([]byte, len(t[0]))
	}

	for y := 0; y < len(t); y++ {
		for x := 0; x < len(t[0]); x++ {
			ret[y][x] = t[x][len(t[0])-y-1]
		}
	}

	return ret
}

func (t Tile) Flip() Tile {
	ret := make(Tile, len(t))
	for i := 0; i < len(t); i++ {
		ret[i] = t[len(t)-i-1]
	}
	return ret
}

func (t Tile) RightSideMatchesLeftOf(neighbour Tile) bool {
	for y := 0; y < len(t); y++ {
		if t[y][len(t[0])-1] != neighbour[y][0] {
			return false
		}
	}
	return true
}

func (t Tile) BottomSideMatchesTopOf(neighbour Tile) bool {
	for x := 0; x < len(t[0]); x++ {
		if t[len(t)-1][x] != neighbour[0][x] {
			return false
		}
	}
	return true
}

func (t Tile) ReplaceSeaMonsters() bool {
	// (y, x) is the top left corner of our sliding window
	var ok bool
	for y := 0; y < len(t)-seaMonsterHeight; y++ {
		for x := 0; x < len(t[0])-seaMonsterWidth; x++ {
			if t.HaveMonsterAt(y, x) {
				t.ReplaceMonsterAt(y, x)
				ok = true
			}
		}
	}
	return ok
}

func (t Tile) HaveMonsterAt(y, x int) bool {
	for monsterY := 0; monsterY < seaMonsterHeight; monsterY++ {
		for monsterX := 0; monsterX < seaMonsterWidth; monsterX++ {
			if seaMonster[monsterY][monsterX] == '#' && t[y+monsterY][x+monsterX] != '#' {
				return false
			}
		}
	}
	return true
}

func (t Tile) ReplaceMonsterAt(y, x int) {
	for monsterY := 0; monsterY < seaMonsterHeight; monsterY++ {
		for monsterX := 0; monsterX < seaMonsterWidth; monsterX++ {
			if seaMonster[monsterY][monsterX] == '#' {
				t[y+monsterY][x+monsterX] = 'O'
			}
		}
	}
}
