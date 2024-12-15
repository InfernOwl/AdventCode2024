package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//Open input file and parse out values
	//Adding values to 2D array to hold them

	//Open the input file
	dat, err := os.Open("../inputs/day4.txt")
	check(err)
	reader := bufio.NewScanner(dat)
	sourceArr := [][]string{}

	//For every line, parse out valid strings with regex
	for reader.Scan() {
		line := reader.Text()
		sourceArr = append(sourceArr, strings.Split(line, ""))
	}

	//fmt.Println(sourceArr[0][len(sourceArr[0])-1])

	//Pass 2D array into partOne and partTwo functions to get solves
	partOne(sourceArr)
}

func partOne(arr [][]string) {
	//Pull 2D array use omnidrectional helper functions to check for 'XMAS'
	//string appearances in array

	//Initialize variables
	patternArr := []string{"X", "M", "A", "S"}
	interFinds := 0
	totalFinds := 0

	//Loop through each entry in array
	//If it matches pattern start, call recursive checkers
	for y := range arr {
		for x := range y {
			if arr[y][x] == patternArr[0] {
				interFinds += checkNorth(arr, patternArr, x, y, 0)
				interFinds += checkSouth(arr, patternArr, x, y, 0)
				interFinds += checkWest(arr, patternArr, x, y, 0)
				interFinds += checkEast(arr, patternArr, x, y, 0)
				interFinds += checkNorthWest(arr, patternArr, x, y, 0)
				interFinds += checkNorthEast(arr, patternArr, x, y, 0)
				interFinds += checkSouthWest(arr, patternArr, x, y, 0)
				interFinds += checkSouthEast(arr, patternArr, x, y, 0)
				fmt.Println(interFinds)
				totalFinds += interFinds
			}
		}
	}

	fmt.Println(interFinds)
}

/*func partTwo(arr [][]string) {

}
*/

//----Omnidirectional recursive 'XMAS' checking functions

//If x and y coordinates(xPos,yPos) in source array(arr) matches letter(let)
//then recurse function in appropriate direction if applicable
//If it is the final letter in goal word, return 1
//else return 0
//(Function is the same for all directions so no further documentation
//will be written)

func checkNorth(arr [][]string, patternArr []string, xPos int, yPos int, let int) int {
	fmt.Println(patternArr[let], arr[yPos][xPos])
	if let < len(patternArr)-1 {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return checkNorth(arr, patternArr, xPos, yPos-1, let+1)
				}
			}
		}
	} else {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					fmt.Println("MADE IT")
					return 1
				}
			}
		}
	}
	return 0
}

func checkSouth(arr [][]string, patternArr []string, xPos int, yPos int, let int) int {
	if let < len(patternArr)-1 {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return checkSouth(arr, patternArr, xPos, yPos+1, let+1)
				}
			}
		}
	} else {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return 1
				}
			}
		}
	}
	return 0
}

func checkWest(arr [][]string, patternArr []string, xPos int, yPos int, let int) int {
	if let < len(patternArr)-1 {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return checkWest(arr, patternArr, xPos-1, yPos, let+1)
				}
			}
		}
	} else {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return 1
				}
			}
		}
	}
	return 0
}

func checkEast(arr [][]string, patternArr []string, xPos int, yPos int, let int) int {
	if let < len(patternArr)-1 {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return checkEast(arr, patternArr, xPos+1, yPos, let+1)
				}
			}
		}
	} else {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return 1
				}
			}
		}
	}
	return 0
}

func checkNorthWest(arr [][]string, patternArr []string, xPos int, yPos int, let int) int {
	if let < len(patternArr)-1 {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return checkNorthWest(arr, patternArr, xPos-1, yPos-1, let+1)
				}
			}
		}
	} else {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return 1
				}
			}
		}
	}
	return 0
}

func checkNorthEast(arr [][]string, patternArr []string, xPos int, yPos int, let int) int {
	if let < len(patternArr)-1 {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return checkNorthEast(arr, patternArr, xPos+1, yPos-1, let+1)
				}
			}
		}
	} else {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return 1
				}
			}
		}
	}
	return 0
}

func checkSouthWest(arr [][]string, patternArr []string, xPos int, yPos int, let int) int {
	if let < len(patternArr)-1 {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return checkSouthWest(arr, patternArr, xPos-1, yPos+1, let+1)
				}
			}
		}
	} else {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return 1
				}
			}
		}
	}
	return 0
}

func checkSouthEast(arr [][]string, patternArr []string, xPos int, yPos int, let int) int {
	if let < len(patternArr)-1 {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return checkSouthEast(arr, patternArr, xPos+1, yPos+1, let+1)
				}
			}
		}
	} else {
		if yPos >= 0 && yPos < len(arr) {
			if xPos >= 0 && xPos < len(arr[yPos]) {
				if arr[yPos][xPos] == patternArr[let] {
					return 1
				}
			}
		}
	}
	return 0
}
