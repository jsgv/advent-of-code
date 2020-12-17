package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var (
	numbers  []int
	preamble = 25
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")
	l := strings.Split(string(b), "\n")

	for _, v := range l {
		n, _ := strconv.Atoi(v)
		numbers = append(numbers, n)
	}

outer:
	for i, n := range numbers {
		if i < preamble {
			continue
		}

		var start = 0

		if i <= preamble {
			start = 0
		} else {
			start = i - preamble
		}

		toTest := numbers[start:i]
		var isValid = false

		for _, d := range toTest {
			for _, d2 := range toTest {
				if d+d2 == n {
					isValid = true
				}
			}
		}

		if !isValid {
			fmt.Println("Part 1:", n)

			for i := range numbers {
				for i2 := i + 1; i2 < len(numbers); i2++ {
					r := numbers[i:i2]
					if sliceSum(r) == n {
						sort.Ints(r)
						fmt.Println("Part 2:", r[0]+r[len(r)-1])
						break outer
					}

				}
			}
		}
	}
}

func sliceSum(numbers []int) int {
	total := 0

	for _, d := range numbers {
		total += d
	}

	return total
}
