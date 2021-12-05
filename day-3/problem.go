package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = "input.txt"
var ASCII = 48

func main() {
	fmt.Print("Part 1\n")
	part1()

	fmt.Print("Part 2\n")
	part2()
}

func part1() {
	var zeros = 0
	var ones = 0
	lines, _ := readLines(input)
	nOfLine := len(lines[0])
	gamma := make([]string, nOfLine)
	epsilon := make([]string, nOfLine)
	split := make([]string, len(lines))
	copy(split, lines)
	fmt.Println(string(split[0][3]))
	for i := 0; i < nOfLine; i++ {
		for line := range split {
			if split[line][i] == byte(0+ASCII) {
				zeros++
			}
			if split[line][i] == byte(1+ASCII) {
				ones++
			}
		}
		if zeros > ones {
			gamma[i] = "0"
			epsilon[i] = "1"
		}
		if zeros < ones {
			gamma[i] = "1"
			epsilon[i] = "0"
		}
		fmt.Println(i, " column most=", gamma[i], "column least=", epsilon[i], "zeros=", zeros, "ones=", ones, "checksum=", zeros+ones)
		zeros = 0
		ones = 0
	}
	gBin := strings.Join(gamma, "")
	eBin := strings.Join(epsilon, "")
	fmt.Println(gBin, eBin)

	gDec, _ := strconv.ParseInt(gBin, 2, 64)
	eDec, _ := strconv.ParseInt(eBin, 2, 64)
	fmt.Println(gDec, eDec, "multiplied output =", gDec*eDec)
}

func part2() {
	lines, _ := readLines(input)
	var currentOxy []string
	copy(currentOxy, lines)
	var currentCo2 []string
	copy(currentCo2, lines)
	oxyRating := int64(0)
	co2Rating := int64(0)
	for i := 0; i < len(lines[0]); i++ {
		var linesOxy []string
		var linesCo2 []string

		for line := range currentOxy {
			if currentOxy[line][i] == byte(1+ASCII) {
				linesOxy = append(linesOxy, currentOxy[line])
			}
		}
		for line := range currentCo2 {
			if currentCo2[line][i] == byte(0+ASCII) {
				linesCo2 = append(linesCo2, currentCo2[line])
			}
		}

		var cOxy []string
		if len(linesOxy) >= len(linesCo2)/2 {
			currentOxy = linesOxy
		} else {
			for currentLine := range currentOxy {
				if !in(currentOxy[currentLine], linesOxy) {
					cOxy = append(cOxy, currentOxy[currentLine])
				}
			}
		}
		var cCo2 []string
		if len(linesOxy) < len(linesCo2)/2 {
			currentCo2 = linesCo2
		} else {
			for currentLine := range currentCo2 {
				if !in(currentOxy[currentLine], linesOxy) {
					cCo2 = append(cCo2, currentCo2[currentLine])
				}
			}
		}

		if len(cOxy) == 1 {
			oxyRating, _ = strconv.ParseInt(strings.Join(cOxy, ""), 2, 64)
		}
		if len(cCo2) == 1 {
			co2Rating, _ = strconv.ParseInt(strings.Join(cCo2, ""), 2, 64)
		}
	}
	fmt.Println(oxyRating, co2Rating)
}

func in(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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
