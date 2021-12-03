package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	items []struct {
		direction string
		value     int
	}
)

func init() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rx := regexp.MustCompile(`(\w+) (\d+)`)

	for scanner.Scan() {
		matches := rx.FindAllStringSubmatch(scanner.Text(), -1)

		for _, m := range matches {
			value, _ := strconv.Atoi(m[2])

			items = append(items, struct {
				direction string
				value     int
			}{
				direction: m[1],
				value:     value,
			})
		}

	}
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	horizontal := 0
	depth := 0

	for _, v := range items {
		switch v.direction {
		case "up":
			depth -= v.value
		case "down":
			depth += v.value
		case "forward":
			horizontal += v.value
		}
	}

	return horizontal * depth
}

func part2() int {
	horizontal := 0
	depth := 0
	aim := 0

	for _, v := range items {
		switch v.direction {
		case "up":
			aim -= v.value
		case "down":
			aim += v.value
		case "forward":
			horizontal += v.value
			depth += v.value * aim
		}
	}

	return horizontal * depth
}
