package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

type line struct {
	row     int
	content string
}

type part struct {
	row        int
	start      int
	end        int
	partNumber int64
	include    bool
}

var lines []line
var parts []part
var sumOfAllPartNumbers int64 = 0
var sumOfAllGearRatios int64 = 0

func part1() {
	fmt.Println("Analyzing schematic...")

	file, err := os.Open("../../inputs/day03/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineNr := 0
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\D+`)

	for scanner.Scan() {
		content := scanner.Text()
		lines = append(lines, line{row: lineNr, content: content})
		split := strings.Split(re.ReplaceAllString(content, " "), " ")
		parsedContent := content

		// This is for part 1
		for _, partNumber := range split {
			if partNumber == "" {
				continue
			}

			index := strings.Index(parsedContent, partNumber)
			partNumberInt, _ := strconv.ParseInt(partNumber, 10, 0)

			part := part{
				row:        lineNr,
				start:      index,
				end:        index + len(partNumber),
				partNumber: partNumberInt,
				include:    false,
			}
			parts = append(parts, part)

			parsedContent = strings.Replace(parsedContent, partNumber, strings.Repeat(".", len(partNumber)), 1)
		}

		lineNr++
	}

	// Now loop over the parts to check if they should be included in the schematics
	for _, part := range parts {
		if (part.row - 1) >= 0 {
			if analyzeRow(part.row-1, part) {
				part.include = true
				sumOfAllPartNumbers += part.partNumber
				continue
			}
		}

		if analyzeRow(part.row, part) {
			part.include = true
			sumOfAllPartNumbers += part.partNumber
			continue
		}

		if (part.row + 1) <= (len(lines) - 1) {
			if analyzeRow(part.row+1, part) {
				part.include = true
				sumOfAllPartNumbers += part.partNumber
				continue
			}
		}
	}

	fmt.Println("Sum of all engine parts:")
	fmt.Println(sumOfAllPartNumbers)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func part2() {
	for index, line := range lines {
		if index == 0 || index == len(lines)-1 {
			continue
		}

		//fmt.Println(line.content)
		gearIndexes := findAllIndexes(line.content, "*")

		for _, gearIndex := range gearIndexes {
			getGearParts(index, gearIndex)
		}
	}

	fmt.Println("Sum of all gear ratios:")
	fmt.Println(sumOfAllGearRatios)
}

func analyzeRow(rowNr int, part part) bool {
	content := lines[rowNr].content

	start := part.start - 1
	if start < 0 {
		start = 0
	}

	end := part.end + 1
	if end > len(content) {
		end = part.end
	}

	substring := content[start:end]

	re := regexp.MustCompile(`[^a-zA-Z0-9.]`)

	return re.MatchString(substring)
}

func findAllIndexes(s, substr string) []int {
	indexes := make([]int, 0)

	// Start searching for the substring from the beginning of the string
	index := -1
	for {
		index = strings.Index(s, substr)

		// If the substring is not found, break the loop
		if index == -1 {
			break
		}

		// Append the index to the slice
		indexes = append(indexes, index)

		// Move the starting point for the next search after the found substring
		s = strings.Replace(s, "*", ".", 1)
	}

	return indexes
}

func getGearParts(rowNr int, checkIndex int) {
	// We only need to check parts in rowNr -1, rowNr and rowNr + 1
	var gearParts []int64

	for _, part := range parts {
		if (part.row == rowNr-1) || (part.row == rowNr) || (part.row == rowNr+1) {
			if (part.start <= checkIndex+1) && ((part.end - 1) >= checkIndex-1) {
				gearParts = append(gearParts, part.partNumber)
			}
		}
	}

	if len(gearParts) == 2 {
		sumOfAllGearRatios += gearParts[0] * gearParts[1]
	}
}
