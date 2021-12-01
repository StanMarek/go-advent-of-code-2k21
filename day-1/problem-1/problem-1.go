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
	0:  "(N/A - no previous measurement)",
	1:  "(increased)",
}

func main() {
	var increased int = 0

	lines, _ := readLines(input)
	nLines := len(lines)
	//fmt.Println(len(lines))
	measurements := make([]int, nLines)

	for i, line := range lines {
		//fmt.Println("%t",line)
		val, _ := strconv.Atoi(line)
		measurements[i] = val
	}

	for i := 0; i < nLines-1; i++ {
		if measurements[i] < measurements[i+1] {
			fmt.Println(measurements[i+1], measurementMap[1])
			increased++
		}
		if measurements[i] == measurements[i+1] {
			fmt.Println(measurements[i+1], measurementMap[0])
		}
		if measurements[i] > measurements[i+1] {
			fmt.Println(measurements[i+1], measurementMap[-1])
		}
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
