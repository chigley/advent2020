package day20_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day20"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 20899048083289, 273},
	{"input", 18482479935793, 2118},
}

func TestDay20(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			groups, err := advent2020.ReadStringGroups(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			tiles, err := day20.ParseTiles(groups)
			if err != nil {
				t.Fatal(err)
			}

			picture, err := day20.BuildPicture(tiles)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day20.Part1(picture))

			part2, err := day20.Part2(picture)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part2, part2)
		})
	}
}
