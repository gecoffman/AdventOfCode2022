package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	totalCalories int
}

var caloriesByElf = []int{}

func main() {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("File where art thou?: %s", err)
	}
	content := string(f)
	lines := strings.Split(content, "\r\n")
	currCal := 0

	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			mealCal, err := strconv.ParseInt(lines[i], 0, 32)
			if err != nil {
				fmt.Println(err)
			}
			currCal += int(mealCal)

		} else {
			caloriesByElf = append(caloriesByElf, currCal)
			currCal = 0
		}
	}
	sortedCalories := caloriesByElf[:]
	sort.Sort(sort.IntSlice(sortedCalories))
	top := sortedCalories[len(sortedCalories)-3:]
	fmt.Printf("The most calories for an elf is: %v\n", top[2])
	fmt.Printf("The 2nd calories for an elf is: %v\n", top[1])
	fmt.Printf("The 3rd calories for an elf is: %v\n", top[0])

	fmt.Printf("The sum of the top 3 calories is: %v\n", top[0]+top[1]+top[2])

}
