// SOLVED
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var input = "../puzzle-input.txt"

var measurementMap = map[int]string{
	-1: "(decreased)",
	0:  "(no diff)",
	1:  "(increased)",
}

func main() {
	var increased int = 0
	sum := 0
	lines, _ := readLines(input)
	nLines := len(lines)
	measurements := make([]int, nLines)
	var sums []int

	for i, line := range lines {
		val, _ := strconv.Atoi(line)
		measurements[i] = val
	}

	for i := 0; i < len(measurements)-2; i++ {
		for j := i; j < i+3; j++ {
			sum += measurements[j]
		}
		fmt.Println(sum)
		sums = append(sums, sum)
		sum = 0
	}

	for i := 0; i < len(sums)-1; i++ {
		if sums[i] < sums[i+1] {
			fmt.Println(sums[i+1], measurementMap[1])
			increased++
		}
		if sums[i] == sums[i+1] {
			fmt.Println(sums[i+1], measurementMap[0])
		}
		if sums[i] > sums[i+1] {
			fmt.Println(sums[i+1], measurementMap[-1])
		}
		//time.Sleep(1000000000)
	}
	fmt.Println(increased)
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
