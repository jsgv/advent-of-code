package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

var (
	items []int
)

func init() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		items = append(items, v)
	}
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	i := 0
	prev := 0

	for _, v := range items {
		if prev == 0 {
			prev = v
			continue
		}

		if v > prev {
			i++
		}

		prev = v
	}

	return i
}

func part2() int {
	i := 0
	prev := 0

	for k := range items {
		v := 0

		for _, v2 := range items[k : k+3] {
			v += v2
		}

		if prev == 0 {
			prev = v
			continue
		}

		if v > prev {
			i++
		}

		prev = v
	}

	return i
}

func TestAPart1(t *testing.T) {
	if part1() != 1602 {
		t.Errorf("Invalid value for part 1 - wanted: [1602] got: [%d]", part1())
	}
}

func TestAPart2(t *testing.T) {
	if part2() != 1633 {
		t.Errorf("Invalid value for part 2 - wanted: [1633] got: [%d]", part2())
	}
}
