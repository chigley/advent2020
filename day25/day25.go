package day25

import (
	"fmt"
)

func Day25(in []int) (int, error) {
	if len(in) != 2 {
		return 0, fmt.Errorf("day25: got %d inputs, expected 2", len(in))
	}
	return Transform(in[0], LoopSize(in[1])), nil
}

func Transform(subject, loopSize int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val *= subject
		val %= 20201227
	}
	return val
}

func LoopSize(publicKey int) int {
	val := 1
	for i := 1; ; i++ {
		val *= 7
		val %= 20201227
		if val == publicKey {
			return i
		}
	}
}
