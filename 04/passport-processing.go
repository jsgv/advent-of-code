package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	hgtRegex       = regexp.MustCompile("^([0-9]{2,3})(cm|in)$")
	hclRegex       = regexp.MustCompile("^#[0-9a-f]{6}$")
	eclRegex       = regexp.MustCompile("amb|blu|brn|gry|grn|hzl|oth")
	pidRegex       = regexp.MustCompile("^[0-9]{9}$")
	requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")

	passports := strings.Split(string(b), "\n\n")
	r := regexp.MustCompile("[ \n]")

	var items []map[string]string

	for _, i := range passports {
		values := make(map[string]string)
		parts := r.Split(i, -1)

		for _, j := range parts {
			if j == "" {
				continue
			}

			v := strings.Split(j, ":")
			values[v[0]] = v[1]
		}

		items = append(items, values)
	}

	// part 1
	totalValidPart1 := 0
	for _, i := range items {
		if !hasRequiredFields(i) {
			continue
		}

		totalValidPart1++
	}
	fmt.Println("Total valid:", totalValidPart1)

	// part 2
	totalValidPart2 := 0
	for _, i := range items {
		if !hasRequiredFields(i) {
			continue
		}

		if !hasValidFields(i) {
			continue
		}

		totalValidPart2++
	}
	fmt.Println("Total valid:", totalValidPart2)
}

func hasRequiredFields(fields map[string]string) bool {
	isValid := true

outer:
	for _, k := range requiredFields {
		if _, ok := fields[k]; !ok {
			isValid = false
			break outer
		}
	}

	return isValid
}

func hasValidFields(fields map[string]string) bool {
	isValid := true

outer:
	for k := range fields {
		value := fields[k]

		if k == "byr" {
			v, _ := strconv.Atoi(value)
			if v < 1920 || v > 2002 {
				isValid = false
			}
		} else if k == "iyr" {
			v, _ := strconv.Atoi(value)
			if v < 2010 || v > 2020 {
				isValid = false
			}
		} else if k == "eyr" {
			v, _ := strconv.Atoi(value)
			if v < 2020 || v > 2030 {
				isValid = false
			}
		} else if k == "hgt" {
			if !hgtRegex.MatchString(value) {
				isValid = false
			} else {
				parts := hgtRegex.FindAllStringSubmatch(value, -1)
				v, _ := strconv.Atoi(parts[0][1])
				if parts[0][2] == "cm" {
					if v < 150 || v > 193 {
						isValid = false
					}
				} else if parts[0][2] == "in" {
					if v < 59 || v > 76 {
						isValid = false
					}
				}
			}
		} else if k == "hcl" {
			if !hclRegex.MatchString(value) {
				isValid = false
			}
		} else if k == "ecl" {
			if !eclRegex.MatchString(value) {
				isValid = false
			}
		} else if k == "pid" {
			if !pidRegex.MatchString(value) {
				isValid = false
			}
		}

		if !isValid {
			break outer
		}
	}

	return isValid
}
