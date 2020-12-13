package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var items []string

	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	// for part 2
	var numbers []int

	// part 1
	max := 0
	for _, i := range items {
		n := 0

		for _, j := range i {
			n <<= 1

			switch j {
			case 'J', 'L':
				n |= 0
			case 'B', 'R':
				n |= 1
			}

		}

		if n > max {
			max = n
		}

		// for part 2
		numbers = append(numbers, n)
	}
	// part 1
	fmt.Println("Highest seat ID:", max)

	// part 2
	sort.Ints(numbers)
	for i := 1; i < len(numbers); i++ {
		if numbers[i]-1 != numbers[i-1] {
			fmt.Println("Missing:", numbers[i]-1)
		}
	}
}
