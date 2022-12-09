package main

import (
	"fmt"
	"strconv"

	"github.com/gecoffman/AdventOfCode2022/lib"
)

type tree struct {
	row       int
	col       int
	height    int
	isVisible bool
	score     int
	scores    []string
}

func getTree(row int, col int, array [][]tree) tree {
	return array[row][col]
}

func (t *tree) setIsVisibe(array [][]tree) {
	rowLen := len(array[0])
	colLen := len(array)
	if t.row == 0 || t.col == 0 || t.col == colLen-1 || t.row == rowLen-1 {
		t.isVisible = true
		t.score = 0
		return
	}
	// look up
	if t.row > 0 {
		currRow := t.row - 1
		allSmaller := true
		score := 0
		for currRow >= 0 {
			score++
			otherT := getTree(currRow, t.col, array)
			if otherT.height >= t.height {
				allSmaller = false
				break
			}
			currRow--
		}
		t.score *= score
		t.scores = append(t.scores, fmt.Sprintf("up score: %v", score))
		if allSmaller {
			t.isVisible = true
		}
	}

	// look down
	if t.row < rowLen-1 {
		currRow := t.row + 1
		allSmaller := true
		score := 0
		for currRow <= rowLen-1 {
			score++
			otherT := getTree(currRow, t.col, array)
			if otherT.height >= t.height {
				allSmaller = false
				break
			}
			currRow++

		}
		t.score *= score
		t.scores = append(t.scores, fmt.Sprintf("down score: %v", score))

		if allSmaller {
			t.isVisible = true
		}

	}
	// look left
	if t.col > 0 {
		currCol := t.col - 1
		allSmaller := true
		score := 0
		for currCol >= 0 {
			score++
			otherT := getTree(t.row, currCol, array)
			if otherT.height >= t.height {
				allSmaller = false
				break
			}
			currCol--
		}
		t.score *= score
		t.scores = append(t.scores, fmt.Sprintf("left score: %v", score))

		if allSmaller {
			t.isVisible = true
		}
	}
	// look right
	if t.col < colLen-1 {
		currCol := t.col + 1
		allSmaller := true
		score := 0
		for currCol <= colLen-1 {
			score++
			otherT := getTree(t.row, currCol, array)
			if otherT.height >= t.height {
				allSmaller = false
				break
			}
			currCol++

		}
		t.score *= score
		t.scores = append(t.scores, fmt.Sprintf("right score: %v", score))

		if allSmaller {
			t.isVisible = true
		}
	}
}

func part1(lines []string) {
	trees := []tree{}
	array := [][]tree{}
	for i, line := range lines {
		array = append(array, []tree{})
		for j, char := range line {
			height, _ := strconv.ParseInt(string(char), 0, 64)
			t := tree{row: i, col: j, height: int(height), isVisible: false, score: 1}
			trees = append(trees, t)
			if j == 0 {
				array[i] = []tree{}
			}
			array[i] = append(array[i], t)
		}
	}
	count := 0
	score := 0
	for _, t := range trees {
		t.setIsVisibe(array)
		if t.isVisible {
			count++
		}
		if t.score > score {
			score = t.score
		}

	}
	fmt.Printf("Part1: Num of visible trees: %v\n", count)
	fmt.Printf("Part2: Score: %v\n", score)

}

func main() {
	lines := lib.GetInputs("./input.txt")
	part1(lines)

}
