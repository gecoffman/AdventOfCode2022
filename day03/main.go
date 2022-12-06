package main

import (
	"fmt"
	"strings"

	"github.com/gecoffman/AdventOfCode2022/lib"
)

func getPriority(letter byte) (priority int) {
	// fmt.Println(string(letter))
	if letter >= 'a' && letter <= 'z' {
		priority = int(letter) % 96
	} else {
		priority = int(letter)%65 + 27
	}
	// fmt.Println(priority)
	return
}

func part1(lines []string) {
	points := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lineLen := len(line)
		compartment1 := line[:(lineLen / 2)]
		compartment2 := line[(lineLen / 2):]
		priority := 0
		for j := 0; j < len(compartment1); j++ {
			letter := compartment1[j]
			if strings.Contains(compartment2, string(letter)) {
				priority = getPriority(letter)
				break
			}
		}
		if priority == 0 {
			fmt.Println(line)
			fmt.Println(compartment1)
			fmt.Println(len(compartment1))
			fmt.Println(compartment2)
			fmt.Println(len(compartment2))
			break
		}
		points += priority
	}
	fmt.Println(points)
}

func part2(lines []string) {
	points := 0
	for i := 0; i < len(lines)/3; i++ {
		elf1 := lines[i*3]
		elf2 := lines[i*3+1]
		elf3 := lines[i*3+2]
		for j := 0; j < len(elf1); j++ {
			letter := elf1[j]
			if strings.Contains(elf2, string(letter)) && strings.Contains(elf3, string(letter)) {
				points += getPriority(letter)
				fmt.Println(string(letter))
				break
			}
		}
	}
	fmt.Println(points)
}

func main() {
	lines := lib.GetInputs("./input.txt")
	part1(lines)
	part2(lines)
}
