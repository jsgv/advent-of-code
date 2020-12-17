package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type I struct {
	operation string
	value     int
	count     int
}

var (
	instructions []*I
	actionRx     = regexp.MustCompile(`([a-z]{3}) ([-|+]?\d+)`)
)

func work(step int, instructions []*I, acc int) int {
	if len(instructions) <= step {
		fmt.Println("Part 2:", acc)
		return acc
	}

	i := instructions[step]

	if i.count > 0 {
		return acc
	}

	i.count += 1

	switch i.operation {
	case "nop":
		return work(step+1, instructions, acc)
	case "acc":
		acc += i.value
		return work(step+1, instructions, acc)
	case "jmp":
		return work(step+i.value, instructions, acc)
	default:
		return acc
	}
}

func main() {
	b, _ := ioutil.ReadFile("input.txt")
	r := strings.Split(string(b), "\n")

	for _, v := range r {
		if v == "" {
			continue
		}

		parts := actionRx.FindStringSubmatch(v)
		value, _ := strconv.Atoi(parts[2])

		instructions = append(instructions, &I{
			operation: parts[1],
			value:     value,
			count:     0,
		})
	}

	fmt.Println("Part 1:", work(0, sliceCopy(), 0))

	for idx, v := range instructions {
		if v.operation == "nop" || v.operation == "jmp" {
			c := sliceCopy()

			switch v.operation {
			case "nop":
				c[idx].operation = "jmp"
				break
			case "jmp":
				c[idx].operation = "nop"
				break
			}

			work(0, c, 0)

		}
	}
}

func sliceCopy() []*I {
	c := make([]*I, len(instructions))

	for i, p := range instructions {
		if p == nil {
			continue
		}

		v := *p

		c[i] = &v
	}

	return c
}
