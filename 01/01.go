package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(content)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()

		newLine := regexp.MustCompile(`[^0-9]+`)

		newerLine := newLine.ReplaceAllString(line, "")

		lines = append(lines, newerLine)
		// fmt.Println(lines)
	}

	numbers := partOne(lines)

	// sum of slice
	var sum int
	for _, num := range numbers {
		sum += num
	}
	fmt.Println(sum)

	defer content.Close()
}

func partOne(lines []string) []int {
	// convert string to int

	var ints []int
	for _, line := range lines {

		lineLength := len(line)

		if lineLength < 2 {

			firstNumber, secondNumber := string(line[0]), string(line[0])

			tmpStr := firstNumber + secondNumber
			fmt.Println(tmpStr)
			tmpInt, _ := strconv.Atoi(tmpStr)
			ints = append(ints, tmpInt)
		} else {

			firstNumber, secondNumber := string(line[0]), string(line[len(line)-1])

			tmpStr := firstNumber + secondNumber
			fmt.Println(tmpStr)
			tmpInt, _ := strconv.Atoi(tmpStr)
			ints = append(ints, tmpInt)

		}

	}

	return ints
}
