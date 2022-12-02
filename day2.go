package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	part1()
	part2()
}

func part1() {

	moves := make(map[string]int)
	moves["X"] = 1
	moves["Y"] = 2
	moves["Z"] = 3

	loses := make(map[string]string)
	loses["A"] = "Y"
	loses["B"] = "Z"
	loses["C"] = "X"

	draws := make(map[string]string)
	draws["A"] = "X"
	draws["B"] = "Y"
	draws["C"] = "Z"

	file, _ := os.ReadFile("./day2.txt")
	var data []string = strings.Split(string(file), "\n")

	score := 0
	for _, v := range data {
		round := strings.Split(v, " ")

		if loses[round[0]] == round[1] {
			//I win
			score += 6

		} else if draws[round[0]] == round[1] {
			//Draw
			score += 3

		}

		score += moves[round[1]]
	}
	println(score)
}

func part2() {

	moves := make(map[string]int)
	moves["A"] = 1 //Rock
	moves["B"] = 2 //Paper
	moves["C"] = 3 //scissors

	beats := make(map[string]string)
	beats["A"] = "B"
	beats["B"] = "C"
	beats["C"] = "A"

	loses := make(map[string]string)
	loses["A"] = "C"
	loses["B"] = "A"
	loses["C"] = "B"

	file, _ := os.ReadFile("./day2.txt")
	var data []string = strings.Split(string(file), "\n")

	count := 0
	for i, v := range data {
		println(i)
		round := strings.Split(v, " ")

		if round[1] == "Z" {
			//Need a win
			myMove := beats[round[0]]
			count += moves[myMove]
			count += 6

		} else if round[1] == "Y" {
			//Need to draw
			myMove := round[0]
			count += moves[myMove]
			count += 3

		} else {
			//Need to lose
			myMove := loses[round[0]]
			count += moves[myMove]
		}
	}
	fmt.Println(count)

}
