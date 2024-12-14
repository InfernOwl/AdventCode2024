package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//Open the input file
	dat, err := os.Open("../inputs/day3.txt")
	check(err)

	fmt.Println("Part Two: ", partTwo(dat))
	fmt.Println("Part One: ", partOne(dat))

}

func partOne(data *os.File) int {
	//Open reader to parse the passed File by line
	//And declare variables
	reader := bufio.NewScanner(data)
	reg1 := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	reg2 := regexp.MustCompile(`[0-9]+`)
	matches := []string{}
	totals := 0

	//For every line, parse out valid strings with regex
	for reader.Scan() {
		line := reader.Text()
		newText := reg1.FindAllString(line, -1)
		matches = append(matches, newText...)
	}

	//For each valid find, add to totals
	for i, val := range matches {
		newVals := reg2.FindAllString(val, -1)
		leftSide, err1 := strconv.Atoi(newVals[0])
		rightSide, err2 := strconv.Atoi(newVals[1])
		check(err1)
		check(err2)
		totals += leftSide*rightSide + (i * 0)
	}

	return totals
}

func partTwo(data *os.File) int {
	//Open reader to parse the passed File by line
	//And declare variables
	reader := bufio.NewScanner(data)
	reg1 := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don\'t\(\)`)
	reg2 := regexp.MustCompile(`[0-9]+`)
	matches := []string{}
	doAdd := true
	totals := 0

	//For every line, parse out valid strings with regex
	for reader.Scan() {
		line := reader.Text()
		newText := reg1.FindAllString(line, -1)
		fmt.Println(newText)
		matches = append(matches, newText...)
	}

	//For each valid find, if currently a "Do" add to totals
	//If not, continue
	for i, val := range matches {
		if val == "do()" {
			doAdd = true
			continue
		} else if val == "don't()" {
			doAdd = false
			continue
		} else {
			if doAdd {
				newVals := reg2.FindAllString(val, -1)
				leftSide, err1 := strconv.Atoi(newVals[0])
				rightSide, err2 := strconv.Atoi(newVals[1])
				check(err1)
				check(err2)
				totals += leftSide*rightSide + (i * 0)
			}
		}
	}

	return totals
}
