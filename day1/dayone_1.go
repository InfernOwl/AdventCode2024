package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//Open the input file
	dat, err := os.Open("../inputs/dayone_1.txt")
	check(err)
	reader := bufio.NewReader(dat)

	//Define the Datasets for each list and distanceTotal
	var lList []int
	var rList []int
	total := 0

	//Loop to read lines and place values in lists
	stopper := "go" //ironic

	for stopper != "stop" {
		line, err := reader.ReadString('\n')
		if err != nil {
			stopper = "stop"
			break
		}

		check(err)
		saniString := strings.Fields(line)
		lStr, err := strconv.Atoi(saniString[0])
		check(err)
		rStr, err := strconv.Atoi(saniString[1])
		check(err)
		lList = append(lList, lStr)
		rList = append(rList, rStr)
	}

	//Sort both lists
	slices.Sort(lList)
	slices.Sort(rList)

	//Loop to grab total distance
	for len(lList) > 0 {
		x := 0
		y := 0

		x, lList = lList[0], lList[1:]
		y, rList = rList[0], rList[1:]

		total = total + int(math.Floor(math.Abs(float64(x-y))))
	}

	fmt.Println(total)
}
