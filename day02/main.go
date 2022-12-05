package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const ROCK = "ROCK"
const PAPER = "PAPER"
const SCISSORS = "SCISSORS"

const WIN = "WIN"
const LOSS = "LOSS"
const DRAW = "DRAW"

var player1Mapping = map[string]string{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var player2Mapping = map[string]string{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

var player2OutcomeMapping = map[string]string{
	"X": LOSS,
	"Y": DRAW,
	"Z": WIN,
}

var pointsMapping = map[string]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,
}

var rulesWinMapping = map[string]string{
	ROCK:     SCISSORS,
	PAPER:    ROCK,
	SCISSORS: PAPER,
}

var rulesLossMapping = map[string]string{
	SCISSORS: ROCK,
	ROCK:     PAPER,
	PAPER:    SCISSORS,
	// How would you go about making this dynamically from the first mapping?
}

func part1(lines []string) {
	points := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		tokens := strings.Split(line, " ")
		player1Input := player1Mapping[tokens[0]]
		player2Input := player2Mapping[tokens[1]]
		points += pointsMapping[player2Input]
		if player1Input == rulesWinMapping[player2Input] {
			// fmt.Println("Player 2 wins")
			points += 6
		} else if player1Input == player2Input {
			// fmt.Println("Draw")
			points += 3
		} else {
			// fmt.Println("Player 1 wins")
		}

	}
	fmt.Println(points)
}

func part2(lines []string) {
	points := 0
	var player2Input string
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		tokens := strings.Split(line, " ")
		player1Input := player1Mapping[tokens[0]]
		player2Outcome := player2OutcomeMapping[tokens[1]]
		if player2Outcome == WIN {
			player2Input = rulesLossMapping[player1Input]
			points += 6
		} else if player2Outcome == LOSS {
			player2Input = rulesWinMapping[player1Input]
			points += 0
		} else {
			player2Input = player1Input
			points += 3
		}
		points += pointsMapping[player2Input]

	}
	fmt.Println(points)

}

func main() {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("File where art thou?: %s", err)
	}
	content := string(f)
	lines := strings.Split(content, "\r\n")
	part1(lines)
	part2(lines)

}
