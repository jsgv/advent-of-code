package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var items []int

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}

		items = append(items, i)
	}

	// part 1
part1:
	for _, i := range items {
		for _, j := range items {
			if i+j == 2020 {
				fmt.Printf("Found: %d * %d = %d\n", i, j, i*j)
				break part1
			}

		}
	}

	// part 2
part2:
	for _, i := range items {
		for _, j := range items {
			for _, k := range items {
				if i+j+k == 2020 {
					fmt.Printf("Found: %d * %d * %d = %d\n", i, j, k, i*j*k)
					break part2
				}
			}
		}
	}
}
