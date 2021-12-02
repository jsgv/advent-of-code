package main

import "testing"

func TestPart1(t *testing.T) {
	if part1() != 1602 {
		t.Errorf("Invalid value for part 1 - wanted: [1602] got: [%d]", part1())
	}
}

func TestPart2(t *testing.T) {
	if part2() != 1633 {
		t.Errorf("Invalid value for part 2 - wanted: [1633] got: [%d]", part2())
	}
}
