package main

import "testing"

func TestPart1(t *testing.T) {
	if part1() != 1868935 {
		t.Errorf("Invalid value for part 1 - wanted: [1868935] got: [%d]", part1())
	}
}

func TestPart2(t *testing.T) {
	if part2() != 1965970888 {
		t.Errorf("Invalid value for part 2 - wanted: [1965970888] got: [%d]", part2())
	}
}
