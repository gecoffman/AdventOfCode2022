package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gecoffman/AdventOfCode2022/lib"
)

type directory struct {
	name string
	size int
}

func getDirectoryName(line string) string {
	tokens := strings.Split(line, " ")
	return tokens[2]
}

func getFileSize(line string) int {
	tokens := strings.Split(line, " ")
	size, _ := strconv.ParseInt(tokens[0], 0, 64)
	return int(size)
}

func print(dirs []*directory) {
	for _, d := range dirs {
		fmt.Printf("\t%s:%v\n", d.name, d.size)
	}
}
func part1(lines []string) ([]*directory, int) {
	directories := []*directory{}
	currDirectories := []*directory{}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "$") {
			if strings.Contains(line, "cd ..") {
				// pop off curr dir
				currDirectories = currDirectories[:len(currDirectories)-1]
			} else if strings.Contains(line, "ls") {

			} else {
				dirName := getDirectoryName(line)
				dir := directory{name: dirName, size: 0}
				currDirectories = append(currDirectories, &dir)
				directories = append(directories, &dir)
			}
		} else {
			if !strings.Contains(line, "dir") {
				size := getFileSize(line)

				for j := 0; j < len(currDirectories); j++ {
					dir := currDirectories[j]
					dir.size += size

				}
			}
		}
	}
	totalSmallSize := 0
	for i := 0; i < len(directories); i++ {
		dir := directories[i]
		if dir.size < 100000 {
			totalSmallSize += dir.size
		}
	}
	return directories, directories[0].size
}

func part2(directories []*directory, totalSize int) {
	fmt.Println(totalSize)
	delSize := 30000000 - (70000000 - totalSize)
	delDir := directories[0]
	fmt.Println(delSize)

	for _, dir := range directories {

		if dir.size >= delSize {
			fmt.Println(dir)
			if delDir.size > dir.size {
				delDir = dir
			}
		}
	}
	fmt.Println(delDir.name)
}

func main() {
	lines := lib.GetInputs("./input.txt")
	directories, totalSize := part1(lines)
	part2(directories, totalSize)
}
