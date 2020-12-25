package day25_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day25"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in  string
	out string
}{
	{"example", "14897079"},
	{"input", "181800"},
}

func TestDay25(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadInts(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			out, err := day25.Day25(in)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.out, out.String())
		})
	}
}

func TestTransform(t *testing.T) {
	assert.Equal(t, "5764801", day25.Transform(7, 8).String())
	assert.Equal(t, "17807724", day25.Transform(7, 11).String())
}

func TestLoopSize(t *testing.T) {
	assert.Equal(t, int64(8), day25.LoopSize(5764801))
	assert.Equal(t, int64(11), day25.LoopSize(17807724))
}
