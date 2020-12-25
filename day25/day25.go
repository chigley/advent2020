package day25

import (
	"fmt"
	"math/big"
)

var modulus = big.NewInt(20201227)

func Day25(in []int) (*big.Int, error) {
	if len(in) != 2 {
		return nil, fmt.Errorf("day25: got %d inputs, expected 2", len(in))
	}

	in1, in2 := int64(in[0]), int64(in[1])
	return Transform(in1, LoopSize(in2)), nil
}

func Transform(subject, loopSize int64) *big.Int {
	var ret big.Int
	return ret.Exp(big.NewInt(subject), big.NewInt(loopSize), modulus)
}

func LoopSize(publicKey int64) int64 {
	val := int64(1)
	for i := int64(1); ; i++ {
		val *= 7
		val %= modulus.Int64()
		if val == publicKey {
			return i
		}
	}
}
