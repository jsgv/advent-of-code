package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Problem struct {
	data []int
}

func sliceToInts(s []string) []int {
	var i []int

	for _, v := range s {
		if v == "" {
			continue
		}

		d, _ := strconv.Atoi(v)
		i = append(i, d)
	}

	sort.Ints(i)

	// add the starting adapter of zero
	i = append([]int{0}, i...)

	// add the last adapter which is largest +3
	i = append(i, i[len(i)-1]+3)

	return i
}

func (p *Problem) part1() int {
	difference1 := 0
	difference3 := 0

	for i, v := range p.data {
		if len(p.data) == i+1 {
			continue
		}

		nextNumber := p.data[i+1]

		if v+1 == nextNumber {
			difference1++
		} else if v+3 == nextNumber {
			difference3++
		}
	}

	return difference1 * difference3
}

func (p *Problem) part2() int {
	memo := make(map[int]int)

	dp(0, p.data, memo)

	return memo[0]
}

func dp(idx int, data []int, memo map[int]int) int {
	if len(data) == 1 {
		return 1
	}

	total := 0

	for _, v := range data {
		if v-idx > 3 {
			break
		}

		if _, ok := memo[v]; !ok {
			memo[v] = dp(v, data[1:], memo)
		}

		total += memo[v]
	}

	return total
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	items := sliceToInts(strings.Split(string(input), "\n"))

	problem := &Problem{
		data: items,
	}

	fmt.Println("Part 1:", problem.part1())
	fmt.Println("Part 2:", problem.part2())
}
