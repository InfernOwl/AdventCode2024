package main

import (
	"bufio"
	"fmt"
	"os"
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
	reader := bufio.NewScanner(dat)

	//Define the maps for each list and similarity score
	lMap := map[string]int{}
	rMap := map[string]int{}
	total := 0

	//Loop to read lines and tally totals in maps
	for reader.Scan() {
		line := reader.Text()
		
		saniString := strings.Fields(line)
		elem, ok := lMap[saniString[0]]

		if (ok) {
			lMap[saniString[0]] = elem + 1
		} else {
			lMap[saniString[0]] = 1
		}

		elem, ok = rMap[saniString[1]]

		if (ok) {
			rMap[saniString[1]] = elem + 1
		} else {
			rMap[saniString[1]] = 1
		}
	}

	//Loop through lMap to get similarity score
	for key, val := range lMap {
		num, err := strconv.Atoi(key)
		check(err)
		total += num * val * rMap[key]
	}

	fmt.Println(total)
}