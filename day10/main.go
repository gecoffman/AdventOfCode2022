package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gecoffman/AdventOfCode2022/lib"
	"golang.org/x/exp/slices"
)

func getInt(l string) int {
	n, _ := strconv.ParseInt(l, 0, 64)
	return int(n)
}

func getSignalStrenths() []int {
	s := 20
	signals := []int{}
	for s <= 220 {
		signals = append(signals, s)
		s += 40
	}
	return signals

}

func part1(lines []string) {
	signals := getSignalStrenths()
	strength := 0
	register_x := 1
	cycle := 0
	for _, line := range lines {
		var incr int
		tokens := strings.Split(line, " ")
		keyword := tokens[0]
		switch keyword {
		case "noop":
			incr = 0
			cycle++
			if slices.Contains(signals, cycle) {
				//fmt.Printf("[%v] %v \n", i, register_x*cycle)
				strength += register_x * cycle
			}
		case "addx":
			incr = getInt(tokens[1])
			cycle++
			if slices.Contains(signals, cycle) {
				//	fmt.Printf("[%v] %v \n", i, register_x*cycle)
				strength += register_x * cycle
			}
			cycle++
			if slices.Contains(signals, cycle) {
				//fmt.Printf("[%v] %v \n", i, register_x*cycle)
				strength += register_x * cycle
			}

		}
		register_x += incr
		//	fmt.Printf("[%v] %v %v %v\n", i, incr, register_x, strength)

	}
	fmt.Println(strength)
}

func printPixie(cycle int, pixieMiddle int) {
	if cycle == pixieMiddle-1 || cycle == pixieMiddle || cycle == pixieMiddle+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}

func printNewLine(cycle int) {
	if cycle != 0 && cycle%40 == 0 {
		fmt.Print("\n")
	}
}

func part2(lines []string) {
	tot_cycles := 240
	i := 0
	pixie_loc := 1
	j := 0
	for i < tot_cycles {
		r := i % 40
		printNewLine(i)
		if j >= len(lines) {
			i++
			continue
		}
		line := lines[j]
		tokens := strings.Split(line, " ")
		keyword := tokens[0]
		switch keyword {
		case "noop":
			printPixie(r, pixie_loc)
		case "addx":
			printPixie(r, pixie_loc)
			i++
			printNewLine(i)
			r := i % 40
			printPixie(r, pixie_loc)
			pixie_loc += getInt(tokens[1])
		}
		i++
		j++
	}

}

func main() {
	lines := lib.GetInputs("./input.txt")
	//part1(lines)
	part2(lines)

}
