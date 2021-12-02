package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	items []int
)

func init() {
	file, _ := os.Open("input.txt")
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
