package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {

	file, _ := os.ReadFile(("day1.txt"))

	var data []string = strings.Split(string(file), "\n")
	elves := []int{0, 0, 0}

	var count int = 0
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			if count > elves[0] {
				elves[0] = count
			}
			count = 0
			sort.Ints(elves)
			continue
		}
		value, _ := strconv.Atoi(data[i])
		count += value
	}
	if count > elves[0] {
		elves[0] = count
	}
	count = 0
	sort.Ints(elves)

	var result int = 0
	for _, v := range elves {
		result += v
	}
	println(result)
}

func part1() {
	file_data, _ := os.ReadFile("day1.txt")

	var data []string = strings.Split(string(file_data), "\n")

	var max, count int = 0, 0
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			if count > max {
				max = count
			}
			count = 0
			continue
		}
		value, _ := strconv.Atoi(data[i])
		count += value
	}
	println(max)
}
