package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	fmt.Println(partOne(dat))

}

func partOne(data *os.File) string {
	//Open reader to parse the passed File by line
	reader := bufio.NewScanner(data)

	//For every line, parse out valiud strings with regex
	for reader.Scan() {
		line := reader.Text()
		reg := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		//matches := []string{}
		newText := reg.FindAllString(line, -1)
		//reg.ReplaceAllString(line, "FOUND")

		fmt.Println(newText)
	}

	return ""
}
