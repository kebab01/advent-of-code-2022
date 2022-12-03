package main

import (
	"os"
	"strings"
)

func main() {

	part1()
	part2()
}

func part2() {

	file, _ := os.ReadFile("day3.txt")

	bags := strings.Split(string(file), "\n")

	score := 0
	for i := 0; i < (len(bags) / 3); i++ {
		//Every third bag is the start of a new group
		bag_index := i * 3
		for _, character := range bags[bag_index] {
			if strings.Contains(bags[bag_index+1], string(character)) && strings.Contains(bags[bag_index+2], string(character)) {
				if character <= 'Z' {
					score += int(character) - 38
				} else if character <= 'z' {
					score += int(character) - 96
				}
				break
			}
		}
	}
	println(score)
}

func part1() {

	file, _ := os.ReadFile("day3.txt")

	bags := strings.Split(string(file), "\n")

	score := 0
	for _, bag := range bags {
		half := len(bag) / 2
		comp1 := bag[:half]
		comp2 := bag[half:]

		for _, character := range comp1 {

			if strings.Contains(comp2, string(character)) {
				if character <= 'Z' {
					score += int(character) - 38
				} else if character <= 'z' {
					score += int(character) - 96
				}
				break
			}

		}
	}
	println(score)

}
