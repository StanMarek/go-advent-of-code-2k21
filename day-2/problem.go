package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = "input.txt"

func main() {
	fmt.Println("PART 1")
	part1()
	fmt.Println("\nPART 2")
	part2()
}

func part1() {
	var horizontal = 0
	var depth = 0
	lines, _ := readLines(input)
	split := make([][]string, len(lines))
	//length := len(lines)
	for i, line := range lines {
		split[i] = strings.Split(line, " ")
	}
	//fmt.Println(split[999][0], split[999][1])

	for i := 0; i < len(lines); i++ {
		if split[i][0] == "forward" {
			val, _ := strconv.Atoi(split[i][1])
			horizontal += val
		}
		if split[i][0] == "down" {
			val, _ := strconv.Atoi(split[i][1])
			depth += val
		}
		if split[i][0] == "up" {
			val, _ := strconv.Atoi(split[i][1])
			depth -= val
		}

	}
	fmt.Println("forward:", horizontal, "depth:", depth)
	fmt.Println("multiplied:", horizontal*depth)
}

func part2() {
	var horizontal = 0
	var depth = 0
	var aim = 0
	lines, _ := readLines(input)
	split := make([][]string, len(lines))
	//length := len(lines)
	for i, line := range lines {
		split[i] = strings.Split(line, " ")
	}
	//fmt.Println(split[999][0], split[999][1])

	for i := 0; i < len(lines); i++ {
		if split[i][0] == "forward" {
			val, _ := strconv.Atoi(split[i][1])
			horizontal += val
			depth += aim * val
		}
		if split[i][0] == "down" {
			val, _ := strconv.Atoi(split[i][1])
			aim += val
		}
		if split[i][0] == "up" {
			val, _ := strconv.Atoi(split[i][1])
			aim -= val
		}

	}

	fmt.Println("forward:", horizontal, "depth:", depth, "aim:", aim)
	fmt.Println("multiplied:", horizontal*depth)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
