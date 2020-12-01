package advent2020

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
	var ret []int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("advent2020: atoi: %w", err)
		}
		ret = append(ret, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("advent2020: scanner: %w", err)
	}

	return ret, nil
}
