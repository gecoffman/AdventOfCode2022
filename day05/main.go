package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gecoffman/AdventOfCode2022/lib"
)

func getStacks(lines []string) ([]string, int) {
	stacks := []string{}
	for i := 0; i < len(lines[0])/3; i++ {
		stacks = append(stacks, "")
	}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		col := 0
		if strings.Contains(line, "[") {
			for j := 0; j < len(line); j += 4 {
				token := line[j : j+3]
				//fmt.Printf("Token [%v] is \"%v\"\n", j, token)
				if len(token) == 3 && token[0] == '[' {
					crate := token[1]
					currStack := stacks[col]
					stacks[col] = currStack + string(crate)

				}
				col++
			}
		} else {
			return stacks, i
		}
	}
	return stacks, 0
}

func moveCrates(lines []string, part string) {
	stacks, row := getStacks(lines)
	fmt.Println(stacks)
	for i := row + 2; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "move") {
			tokens := strings.Split(line, " ")
			target, _ := strconv.ParseInt(tokens[5], 0, 64)
			src, _ := strconv.ParseInt(tokens[3], 0, 64)
			count, _ := strconv.ParseInt(tokens[1], 0, 64)
			//fmt.Printf("Source: %v, target:%v, count:%v\n", src, target, count)
			srcStack := stacks[src-1]
			//targetStack := stacks[target-1]
			crates := srcStack[0:count]
			// fmt.Printf("Crates:%s, src:%s, target:%s\n", crates, srcStack, targetStack)
			switch part {
			case "part1":
				{
					for j := 0; j < len(crates); j++ {
						stacks[target-1] = string(crates[j]) + stacks[target-1]
					}
				}
			case "part2":
				{
					stacks[target-1] = crates + stacks[target-1]
				}
			}

			stacks[src-1] = srcStack[count:]
			//fmt.Println(stacks)
		}
	}
	//fmt.Println(stacks)
	for i := 0; i < len(stacks); i++ {
		if len(stacks[i]) > 0 {
			fmt.Print(string(stacks[i][0]))
		}

	}
	fmt.Print("\n")
}

func main() {
	lines := lib.GetInputs("./input.txt")
	moveCrates(lines, "part1")
	moveCrates(lines, "part2")
}
