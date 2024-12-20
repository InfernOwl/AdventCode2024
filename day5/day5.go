package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	data, err := os.Open("../inputs/day5.txt")
	check(err)
	reader := bufio.NewScanner(data)

	//Declare variables
	rules := map[string][]string{}
	total := 0

	//Read each line of the source input
	//3 options
	// - Line contains '|' symbol = Run function to add to criteria
	// - Blank Line = Ignored (Meant to split criteria and updates)
	// - Line does not contain '|' symbol = List of updates. Must be checked
	//		against criteria

	for reader.Scan() {
		line := reader.Text()

		if line != "" {
			if strings.Contains(line, "|") {
				addRule(rules, line)
			} else {
				//Part One
				total += partOne(rules, line)

				//Part Two
				// total += partTwo(rules, line, false)
			}
		}
	}

	fmt.Println(total)
}

func addRule(rules map[string][]string, line string) {
	//Take line and split between '|'
	//Check rules for digit on the left.
	//If it doesn't exist, add rul with array consiting of digit on
	//	the right to rules
	//If it does exist, append digit on the right to existing rule

	newRule := strings.Split(line, "|")

	_, ok := rules[newRule[0]]
	if ok {
		rules[newRule[0]] = append(rules[newRule[0]], newRule[1])
	} else {
		rules[newRule[0]] = []string{newRule[1]}
	}
}

func partOne(rules map[string][]string, update string) int {
	//Take the update and split into an array of page numbers
	//For each rule, find the first location of the pivot page in the update
	//If not first entry, join every page before pivot page into one string
	//Join rule parameters split by '|' and use to create Regex to check against
	updArr := strings.Split(update, ",")

	for pivot, pages := range rules {
		index := slices.Index(updArr, pivot)
		if index != -1 && index != 0 {
			updString := strings.Join(updArr[:index], ",")
			pagesString := strings.Join(pages, "|")
			ruleRegex := regexp.MustCompile(pagesString)

			ruleCheck := ruleRegex.FindAllString(updString, -1)

			//If it fails the rule check, return nothing
			if ruleCheck != nil {
				return 0
			}
		}
	}

	// If rule check does not fail. Return the center page
	answer, err := strconv.Atoi(updArr[len(updArr)/2])
	check(err)

	return answer

}

func partTwo(rules map[string][]string, update string, isFailure bool) int {
	//Take the update and split into an array of page numbers
	//For each rule, find the first location of the pivot page in the update
	//If not first entry, join every page before pivot page into one string
	//Join rule parameters split by '|' and use to create Regex to check against
	updArr := strings.Split(update, ",")

	for pivot, pages := range rules {
		index := slices.Index(updArr, pivot)
		if index != -1 && index != 0 {
			updString := strings.Join(updArr[:index], ",")
			pagesString := strings.Join(pages, "|")
			ruleRegex := regexp.MustCompile(pagesString)

			ruleCheck := ruleRegex.FindAllString(updString, -1)

			//If it fails the rule check
			//Remove fail numbers from beginning of string
			//Add those numbers (ruleCheck result) after pivot value
			//Append rest of array
			//Recursively run partTwo to check if new update is valid
			if ruleCheck != nil {
				beginning := strings.Split(updString, ",")

				for _, num := range ruleCheck {
					numIndex := slices.Index(beginning, num)
					if numIndex != -1 {
						beginning = append(beginning[:numIndex], beginning[numIndex+1:]...)
					}
				}
				newUpdArr := append(append(append(beginning, updArr[index]), ruleCheck...), updArr[index+1:]...)
				newUpdate := strings.Join(newUpdArr, ",")
				return partTwo(rules, newUpdate, true)
			}
		}
	}

	// If rule check does not fail. Return the center page ONLY if it had failed before
	answer := 0
	if isFailure {
		answer, _ = strconv.Atoi(updArr[len(updArr)/2])
	} else {
		answer = 0
	}

	return answer

}
