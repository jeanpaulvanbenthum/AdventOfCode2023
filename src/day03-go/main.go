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
	//part2()
}

/*
467..114..
...*......
..35..633.
......#...


per part:

r-1 = regel erboven
r=0 = huidige regel part
r1 = regel eronder

r-1: substr: part start -1 tot part end + 1.. zit daar een symbol in? --> include
r-0: is char oo start -1 of end + 1 een symbol? --? include
r-1: substr: part start -1 tot part end + 1.. zit daar een symbol in? --> include
*/

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

func part1() {
	fmt.Println("Analyzing schematic...")

	file, err := os.Open("../../inputs/day03/input.txt")
	//file, err := os.Open("../../inputs/day03/example-1.txt")
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

	fmt.Println(sumOfAllPartNumbers)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func analyzeRow(rowNr int, part part) bool {
	content := lines[rowNr].content
	fmt.Println(part)

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

func part2() {

}
