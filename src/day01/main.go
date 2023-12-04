package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
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
