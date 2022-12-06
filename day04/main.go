package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gecoffman/AdventOfCode2022/lib"
)

type CleaningArea struct {
	start int
	end   int
}

func getCleaningArea(token string) (cleaningArea CleaningArea) {
	nums := strings.Split(token, "-")
	x, _ := strconv.ParseInt(nums[0], 0, 64)
	y, _ := strconv.ParseInt(nums[1], 0, 64)
	cleaningArea = CleaningArea{start: int(x), end: int(y)}
	// fmt.Println(cleaningArea)
	return
}

func (cleaningAreaA *CleaningArea) contains(cleaningAreaB CleaningArea) bool {
	if cleaningAreaA.start <= cleaningAreaB.start {
		if cleaningAreaA.end >= cleaningAreaB.end {
			return true
		}
	}
	return false
}

func (cleaningAreaA *CleaningArea) overlap(cleaningAreaB CleaningArea) bool {
	if cleaningAreaA.end < cleaningAreaB.start {
		return false
	}
	if cleaningAreaA.start > cleaningAreaB.end {
		return false
	}
	return true
}

func part1(lines []string) {
	count := 0
	for i := 0; i < len(lines); i++ {
		tokens := strings.Split(lines[i], ",")
		cleaningArea1 := getCleaningArea(tokens[0])
		cleaningArea2 := getCleaningArea(tokens[1])
		if cleaningArea1.contains(cleaningArea2) {
			count++
		} else if cleaningArea2.contains(cleaningArea1) {
			count++
		}
	}
	fmt.Println(count)
}

func part2(lines []string) {
	count := 0
	for i := 0; i < len(lines); i++ {
		tokens := strings.Split(lines[i], ",")
		cleaningArea1 := getCleaningArea(tokens[0])
		cleaningArea2 := getCleaningArea(tokens[1])
		if cleaningArea1.overlap(cleaningArea2) {
			count++
		} else if cleaningArea2.overlap(cleaningArea1) {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	lines := lib.GetInputs("./input.txt")
	part1(lines)
	part2(lines)
}
