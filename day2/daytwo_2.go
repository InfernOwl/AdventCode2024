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
	return dist > 0 && dist < 4
}

func removeIndex(arr []string, pos int) []string {
	newArr := []string{}
	k := 0
	for i := 0; i < (len(arr)); {
		if i != pos {
			newArr = append(newArr, arr[k])
			k++
		} else {
			k++
		}
		i++
	}

	return newArr
}

func isReportSafe(report []string) bool {

	isAsc := false  //Ascending check
	isDesc := false //Descending check

	for index := range len(report) - 1 {
		curr, currErr := strconv.Atoi(report[index])
		next, nextErr := strconv.Atoi(report[index+1])

		check(currErr)
		check(nextErr)

		// Check adjDifference
		if !adjDifference(curr, next) {
			return false
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
				return false
			}
		} else if isAsc {
			if curr > next {
				return false
			}
		} else if isDesc {
			if curr < next {
				return false
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

		for index := range len(saniString) {
			mixed := removeIndex(saniString, index)
			if isReportSafe(mixed) {
				fmt.Println(saniString, index)
				fmt.Println(mixed)
				safety++ //If report is safe increase tracker
				break
			}
		}
	}

	fmt.Println(safety)
}
