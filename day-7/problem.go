package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}

func part1() {
	input := readFile()

	max, min := max_min(input)
	fmt.Println(max, min)
	bestPosition := min
	var minFuel int
	for i, pos := range input {
		minFuel += pos * i
	}
	for position := min; position <= max; position++ {
		currentFuel := 0
		for _, line := range input {
			fuel := math.Abs(float64(position - line))
			//fmt.Println(fuel)
			currentFuel += int(fuel)
		}
		if currentFuel < minFuel {
			minFuel = currentFuel
			bestPosition = position
		}
	}
	fmt.Println("Best position", bestPosition, "Fuel", minFuel)
}

func part2() {
	input := readFile()

	max, min := max_min(input)
	fmt.Println(max, min)
	bestPosition := min
	var minFuel int
	for i, pos := range input {
		minFuel += pos * i
	}
	for position := min; position <= max; position++ {
		currentFuel := 0
		for _, line := range input {
			crabFuel := 0
			fuel := math.Abs(float64(position - line))
			//fmt.Println(fuel)
			for step := 1; step <= int(fuel); step++ {
				crabFuel += step
			}
			currentFuel += int(crabFuel)
		}
		if currentFuel < minFuel {
			minFuel = currentFuel
			bestPosition = position
		}
	}
	fmt.Println("Best position", bestPosition, "Fuel", minFuel)
}

func readFile() []int {
	input := "input.txt"
	body, _ := ioutil.ReadFile(input)
	var data []int
	content := strings.Split(string(body), ",")

	for _, line := range content {
		val, _ := strconv.Atoi(line)
		//fmt.Println(val)
		data = append(data, val)
	}

	return data
}

func max_min(input []int) (int, int) {
	var max int = input[0]
	var min int = input[0]

	for _, val := range input {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}

	return max, min
}
