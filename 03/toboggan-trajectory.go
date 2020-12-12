package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	slope1 := countTrees(1, 1, lines)
	slope2 := countTrees(3, 1, lines)
	slope3 := countTrees(5, 1, lines)
	slope4 := countTrees(7, 1, lines)
	slope5 := countTrees(1, 2, lines)

	// part 1
	fmt.Println("Total trees:", slope2)

	// part 2
	totalAll := slope1 * slope2 * slope3 * slope4 * slope5
	fmt.Println("Total trees:", totalAll)
}

func countTrees(right int, down int, lines []string) int {
	totalTrees := 0

	for i, l := range lines {
		if i == 0 {
			continue
		} else if down != 1 && i%down != 0 {
			continue
		}

		target := (i / down) * right % len(l)

		if l[target] == '#' {
			totalTrees++
		}
	}

	return totalTrees
}
