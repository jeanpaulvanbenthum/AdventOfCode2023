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
	part2()
}

func part2() {
	var replacer = strings.NewReplacer(
		"one", "1ne",
		"two", "2wo",
		"three", "3hree",
		"four", "4our",
		"five", "5ive",
		"six", "6ix",
		"seven", "7even",
		"eight", "8ight",
		"nine", "9ine",
	)

	fmt.Println("Reading your calibration document:")

	file, err := os.Open("../../inputs/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var calibrationSum int64
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\D+`)

	for scanner.Scan() {
		line := scanner.Text()

		for line != replacer.Replace(line) {
			line = replacer.Replace(line)
		}
		stripped := re.ReplaceAllString(line, "")

		firstInt := stripped[0:1]
		lastInt := stripped[len(stripped)-1:]

		calibrationValue, _ := strconv.ParseInt(firstInt+lastInt, 10, 0)
		calibrationSum += calibrationValue

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Calibration sum: " + strconv.FormatInt(calibrationSum, 10))
}

func part1() {
	fmt.Println("Paste your calibration document:")
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()

		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	var calibrationSum int64

	for index, eachLine := range lines {
		re := regexp.MustCompile(`\D+`)

		stripped := re.ReplaceAllString(eachLine, "")
		firstInt := stripped[0:1]
		lastInt := stripped[len(stripped)-1:]
		calibrationValue, _ := strconv.ParseInt(firstInt+lastInt, 10, 0)
		calibrationSum += calibrationValue
		fmt.Println("Calibration value line " + strconv.Itoa(index) + ": " + strconv.FormatInt(calibrationValue, 10))
	}
	fmt.Println("Calibration sum: " + strconv.FormatInt(calibrationSum, 10))
}
