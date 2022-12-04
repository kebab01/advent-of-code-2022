package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {

	file, _ := os.ReadFile("day4.txt")

	data := strings.Split(string(file), "\n")

	var count int
	for _, item := range data {
		pairs := strings.Split(item, ",")
		p1 := pairs[0]
		p2 := pairs[1]

		p1_lower, _ := strconv.Atoi(strings.Split(p1, "-")[0])
		p1_upper, _ := strconv.Atoi(strings.Split(p1, "-")[1])

		p2_lower, _ := strconv.Atoi(strings.Split(p2, "-")[0])
		p2_upper, _ := strconv.Atoi(strings.Split(p2, "-")[1])

		if p1_upper < p2_lower || p1_lower > p2_upper {
			//No overlap at all
			continue
		}
		count++
	}
	println(count)
}

func part1() {

	file, _ := os.ReadFile("day4.txt")

	data := strings.Split(string(file), "\n")

	var count int
	for _, item := range data {
		pairs := strings.Split(item, ",")
		p1 := pairs[0]
		p2 := pairs[1]

		p1_lower, _ := strconv.Atoi(strings.Split(p1, "-")[0])
		p1_upper, _ := strconv.Atoi(strings.Split(p1, "-")[1])

		p2_lower, _ := strconv.Atoi(strings.Split(p2, "-")[0])
		p2_upper, _ := strconv.Atoi(strings.Split(p2, "-")[1])

		if p1_lower <= p2_lower && p1_upper >= p2_upper {
			// p2 fully contained within p1
			count += 1

		} else if p2_lower <= p1_lower && p2_upper >= p1_upper {
			//P1 fully contained within p2
			count += 1
		}

	}
	fmt.Println(count)
}
