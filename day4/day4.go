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

	//Pass 2D array into partOne and partTwo functions to get solves
	//partOne(sourceArr)
	partTwo(sourceArr)
}

func partOne(arr [][]string) {
	//Pull 2D array use omnidrectional helper functions to check for 'XMAS'
	//string appearances in array

	//Initialize variables
	patternArr := []string{"X", "M", "A", "S"}
	totalFinds := 0

	//Loop through each entry in array
	//If it matches pattern start, call recursive checkers
	for y := range arr {
		for x := range arr {
			if arr[y][x] == patternArr[0] {
				totalFinds += checkNorth(arr, patternArr, x, y, 0)
				totalFinds += checkSouth(arr, patternArr, x, y, 0)
				totalFinds += checkWest(arr, patternArr, x, y, 0)
				totalFinds += checkEast(arr, patternArr, x, y, 0)
				totalFinds += checkNorthWest(arr, patternArr, x, y, 0)
				totalFinds += checkNorthEast(arr, patternArr, x, y, 0)
				totalFinds += checkSouthWest(arr, patternArr, x, y, 0)
				totalFinds += checkSouthEast(arr, patternArr, x, y, 0)
			}
		}
	}

	fmt.Println(totalFinds)
}

func partTwo(arr [][]string) {
	//Pull 2D array use omnidrectional helper functions to check for 'MAS' and 'SAM'
	//string appearances in array oriented around 'A'

	//Initialize variables
	totalFinds := 0

	//Loop through each entry in array
	//If it matches pattern start, call recursive checkers
	for y := range arr {
		for x := range arr {
			if arr[y][x] == "A" {
				if leftDiagMas(arr, x, y) && rightDiagMas(arr, x, y) {
					totalFinds += 1
				}
			}
		}
	}

	fmt.Println(totalFinds)

}

//----Omnidirectional recursive 'XMAS' checking functions

//If x and y coordinates(xPos,yPos) in source array(arr) matches letter(let)
//then recurse function in appropriate direction if applicable
//If it is the final letter in goal word, return 1
//else return 0
//(Function is the same for all directions so no further documentation
//will be written)

func checkNorth(arr [][]string, patternArr []string, xPos int, yPos int, let int) int {
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
					fmt.Println("N", "Success")
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
					fmt.Println("S", "Success")
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
					fmt.Println("W", "Success")
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
					fmt.Println("E", "Success")
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
					fmt.Println("NW", "Success")
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
					fmt.Println("NE", "Success")
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
					fmt.Println("SW", "Success")
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
					fmt.Println("SE", "Success")
					return 1
				}
			}
		}
	}
	return 0
}

func leftDiagMas(arr [][]string, x int, y int) bool {
	//Concat top left, center, and bottom right values if within bounds
	//Compare to 'MAS' and 'SAM'
	//Return true if they match
	phrase := ""

	leftBound := x-1 >= 0 && x-1 < len(arr)
	rightBound := x+1 >= 0 && x+1 < len(arr)
	topBound := y-1 >= 0 && y-1 < len(arr)
	bottomBound := y+1 >= 0 && y+1 < len(arr)

	if leftBound && topBound {
		phrase = phrase + arr[y-1][x-1]
	}

	phrase = phrase + arr[y][x]

	if rightBound && bottomBound {
		phrase = phrase + arr[y+1][x+1]
	}

	return phrase == "MAS" || phrase == "SAM"
}

func rightDiagMas(arr [][]string, x int, y int) bool {
	phrase := ""

	leftBound := x-1 >= 0 && x-1 < len(arr)
	rightBound := x+1 >= 0 && x+1 < len(arr)
	topBound := y-1 >= 0 && y-1 < len(arr)
	bottomBound := y+1 >= 0 && y+1 < len(arr)

	if rightBound && topBound {
		phrase = phrase + arr[y-1][x+1]
	}

	phrase = phrase + arr[y][x]

	if leftBound && bottomBound {
		phrase = phrase + arr[y+1][x-1]
	}

	return phrase == "MAS" || phrase == "SAM"
}
