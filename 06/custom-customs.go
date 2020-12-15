package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")
	groups := strings.Split(string(b), "\n\n")

	// part 1
	totalUniquePart1 := 0
	for _, g := range groups {
		p := strings.Replace(g, "\n", "", -1)
		unique := ""

		for _, c := range p {
			if !strings.Contains(unique, string(c)) {
				unique += string(c)
			}
		}

		totalUniquePart1 += len(unique)
	}
	fmt.Println("Total:", totalUniquePart1)

	totalUniquePart2 := 0
	for _, g := range groups {
		p := strings.Split(g, "\n")
		unique := make(map[int32]int)

		groupAdds := 0
		for _, u := range p {
			for _, c := range u {
				unique[c]++

				if unique[c] == len(p) {
					groupAdds++
					totalUniquePart2++
				}
			}
		}
	}

	fmt.Println("Total:", totalUniquePart2)
}
