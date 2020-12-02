package advent2020

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadIntsFile(name string) ([]int, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("advent2020: opening file: %w", err)
	}
	defer f.Close()

	return readInts(f)
}

func readInts(r io.Reader) ([]int, error) {
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

func ReadStringsFile(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("advent2020: opening file: %w", err)
	}
	defer f.Close()

	return readStrings(f)
}

func readStrings(r io.Reader) ([]string, error) {
	var ret []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("advent2020: scanner: %w", err)
	}

	return ret, nil
}
