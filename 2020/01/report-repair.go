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
		i, _ := strconv.Atoi(scanner.Text())
		items = append(items, i)
	}

part1:
	for _, i := range items {
		for _, j := range items {
			if i+j == 2020 {
				fmt.Printf("Found: %d * %d = %d\n", i, j, i*j)
				break part1
			}
		}
	}

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
