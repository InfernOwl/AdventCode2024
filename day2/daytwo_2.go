package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func adjDifference(x int, y int) bool {
	//Return whether two values are more than
	// 0 and less than 4 away from each other

	dist := int(math.Abs(float64(x - y)))
	return dist < 4
}

func isReportSafe(report []string, secondChance bool) bool {

	isAsc := false  //Ascending check
	isDesc := false //Descending check

	for index := range len(report) - 1 {

		curr, currErr := strconv.Atoi(report[index])
		next, nextErr := strconv.Atoi(report[index+1])

		check(currErr)
		check(nextErr)

		// Check adjDifference
		if !adjDifference(curr, next) {
			fmt.Println(report, secondChance)
			if secondChance {
				report = append(report[:index+1], report[index+2:]...)
				return isReportSafe(report, false)
			} else {
				return false
			}
		}

		// Determine Asc or Desc
		// If already one or the other, skip to safety check
		if !isAsc && !isDesc {
			if curr > next {
				isDesc = true
			} else if curr < next {
				isAsc = true
			} else {
				// If neither greater nor less than
				// Report is not safe is this is the second time

				fmt.Println(report, secondChance)
				if secondChance {
					report = append(report[:index+1], report[index+2:]...)
					return isReportSafe(report, false)
				} else {
					return false
				}

			}
		} else if isAsc {
			if curr > next {

				fmt.Println(report, secondChance)
				if secondChance {
					report = append(report[:index+1], report[index+2:]...)
					return isReportSafe(report, false)
				} else {
					return false
				}
			}
		} else if isDesc {
			if curr < next {
				fmt.Println(report, secondChance)
				if secondChance {
					report = append(report[:index+1], report[index+2:]...)
					return isReportSafe(report, false)
				} else {
					return false
				}
			}
		}
	}

	// If we have not returned false yet, it is safe
	return true
}

func main() {
	//Open the input file
	dat, err := os.Open("../inputs/daytwo.txt")
	check(err)
	reader := bufio.NewScanner(dat)

	//Set variables
	safety := 0 //Safety totals

	//Loop to read lines and determine report safety
	for reader.Scan() {
		line := reader.Text()
		saniString := strings.Fields(line)

		if isReportSafe(saniString, true) {
			safety++ //If report is safe increase tracker
		}
	}

	fmt.Println(safety)
}
