package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Password struct {
	ValidRange []int
	Desired    string
	Value      string
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var items []Password

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		minValue, _ := strconv.Atoi(strings.Split(parts[0], "-")[0])
		maxValue, _ := strconv.Atoi(strings.Split(parts[0], "-")[1])

		items = append(items, Password{
			ValidRange: []int{
				minValue,
				maxValue,
			},
			Desired: strings.TrimSuffix(parts[1], ":"),
			Value:   parts[2],
		})
	}

	// part 1
	totalValidPart1 := 0
	for _, p := range items {
		count := strings.Count(p.Value, p.Desired)
		if p.ValidRange[0] <= count && p.ValidRange[1] >= count {
			totalValidPart1++
		}
	}

	fmt.Printf("Part 1 Total valid: %d\n", totalValidPart1)

	// part 2
	totalValidPart2 := 0
	for _, p := range items {
		if (string(p.Value[p.ValidRange[0]-1]) == p.Desired) !=
			(string(p.Value[p.ValidRange[1]-1]) == p.Desired) {
			totalValidPart2++
		}

	}

	fmt.Printf("Part 2 Total valid: %d\n", totalValidPart2)
}
