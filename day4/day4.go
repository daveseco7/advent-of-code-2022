package day4

import (
	"strings"

	"github.com/daveseco7/advent-of-code-2022/util"
)

type section struct {
	start int
	end   int
}

type pair struct {
	elf1 section
	elf2 section
}

func (p pair) fullyIntersect() bool {
	if p.elf1.start >= p.elf2.start && p.elf1.end <= p.elf2.end {
		// section 1 fully contained in section 2
		return true
	}

	if p.elf2.start >= p.elf1.start && p.elf2.end <= p.elf1.end {
		// section 2 fully contained in section 1
		return true
	}

	return false
}

func (p pair) intersect() bool {
	if p.elf1.end < p.elf2.start {
		return false
	}

	if p.elf1.start > p.elf2.end {
		return false
	}

	return true
}

type area []pair

func parseSection(str string) section {
	strs := strings.Split(str, "-")
	return section{
		start: util.MustAtoi(strs[0]),
		end:   util.MustAtoi(strs[1]),
	}
}

func parseLines(lines []string) area {
	a := make(area, 0, len(lines))
	for _, l := range lines {
		strs := strings.Split(l, ",")

		s1 := parseSection(strs[0])
		s2 := parseSection(strs[1])

		p := pair{
			elf1: s1,
			elf2: s2,
		}

		a = append(a, p)
	}

	return a
}

func Exe1(lines []string) int {
	areaToEvaluate := parseLines(lines)

	intersections := 0
	for _, p := range areaToEvaluate {
		if p.fullyIntersect() {
			intersections++
		}
	}

	return intersections
}

func Exe2(lines []string) int {
	areaToEvaluate := parseLines(lines)

	intersections := 0
	for _, p := range areaToEvaluate {
		if p.intersect() {
			intersections++
		}
	}

	return intersections
}
