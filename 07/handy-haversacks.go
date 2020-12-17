package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	name string
}

var (
	initialBag   = "shiny gold"
	bagNameRx    = regexp.MustCompile(`(.*) bags contain`)
	bagContainRx = regexp.MustCompile(`(\d+?) (.+?) bags?`)
)

func sliceContains(s []string, needle string) bool {
	for _, i := range s {
		if i == needle {
			return true
		}
	}

	return false
}

func findBags(needles []string, items []string, seen map[string]struct{}) int {
	var bags []string

	for _, i := range items {
		if i == "" {
			continue
		}

		holdingBagName := bagNameRx.FindStringSubmatch(i)[1]
		childrenBags := bagContainRx.FindAllStringSubmatch(i, -1)

		for _, c := range childrenBags {
			childBagName := c[2]

			if sliceContains(needles, childBagName) {
				seen[holdingBagName] = struct{}{}
				bags = append(bags, holdingBagName)
			}
		}
	}

	if len(bags) > 0 {
		return findBags(bags, items, seen)
	} else {
		return len(seen)
	}
}

func findBags2(needle string, items []string) int {
	count := 0

	for _, i := range items {
		if i == "" {
			continue
		}

		holdingBagName := bagNameRx.FindStringSubmatch(i)[1]
		childrenBags := bagContainRx.FindAllStringSubmatch(i, -1)

		for _, c := range childrenBags {
			childBagCount, _ := strconv.Atoi(c[1])
			childBagName := c[2]

			if needle == holdingBagName {
				count += childBagCount
				insideCount := findBags2(childBagName, items)
				count += insideCount * childBagCount
			}
		}
	}

	return count
}

func main() {
	b, _ := ioutil.ReadFile("input.txt")
	items := strings.Split(string(b), "\n")

	totalPart1 := findBags([]string{"shiny gold"}, items, make(map[string]struct{}))
	fmt.Println("Total part 1:", totalPart1)

	totalPart2 := findBags2("shiny gold", items)
	fmt.Println("Total part 2:", totalPart2)
}
