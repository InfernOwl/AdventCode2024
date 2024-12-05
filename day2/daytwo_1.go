package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//Open the input file
	dat, err := os.Open("../inputs/daytwo.txt")
	check(err)
	reader := bufio.NewScanner(dat)

	//Set variables
	/*isAsc := false //Ascending check
	isDesc := false //Descending check
	safety := 0 //Safety totals*/
	lastVal := nil //Last value of report
	currVal := 0 //Current value of report

	//Loop to read lines and determine report safety
	for reader.Scan() {
		line := reader.Text()
		saniString := strings.Fields(line)

		for item := range saniString {
			// 
		}
	}
}