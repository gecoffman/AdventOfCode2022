package main

import (
	"fmt"
	"strings"

	"github.com/gecoffman/AdventOfCode2022/lib"
)

func parse(lines []string, packetLen int) {
	line := lines[0]
	marker := ""
	for i := 0; i < len(line); i++ {
		char := string(line[i])
		if !strings.Contains(marker, char) {
			marker = marker + char
		} else {
			x := strings.Index(marker, char)
			marker = marker[x+1:] + char
		}
		//fmt.Printf("[%v] %s %s\n", i, char, marker)
		if len(marker) == packetLen {
			fmt.Println(marker)
			fmt.Println(i + 1)
			break
		}
	}
}

func part1(lines []string) {
	parse(lines, 4)
}

func part2(lines []string) {
	parse(lines, 14)
}

func main() {
	lines := lib.GetInputs("./input.txt")
	part1(lines)
	part2(lines)
}
