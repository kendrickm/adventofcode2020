package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var pageLength, pageWidth int

func parseInput() []string {
	dat, err := ioutil.ReadFile("day3.input")
	check(err)

	page := strings.Split(string(dat), "\n")

	pageLength = len(page)
	pageWidth = len(page[0])

	return page
}

func calculateTrip(rightInc, downInc int, page []string) int {
	tobY := 0
	treeCount := 0

	for tobX := 0; tobX < pageLength; tobX += downInc {
		if (tobX == 0) || (len(page[tobX]) == 0) {
			fmt.Println("Skip")
		} else {

			if tobY >= pageWidth {
				tobY -= pageWidth
			}
			fmt.Printf("Location of sled is %d and %d", tobY, tobX)
			fmt.Println()
			cell := string([]byte{page[tobX][tobY]})
			if cell == "#" {
				fmt.Println("Found tree")
				treeCount++
			}
		}
		tobY += rightInc
	}
	return treeCount
}

func main() {
	page := parseInput()

	firstPart := calculateTrip(3, 1, page)
	fmt.Println(firstPart)

	secondPart := calculateTrip(1, 1, page)
	fmt.Println(secondPart)

	thirdPart := calculateTrip(5, 1, page)
	fmt.Println(thirdPart)

	fourthPart := calculateTrip(7, 1, page)
	fmt.Println(fourthPart)

	fifthPart := calculateTrip(1, 2, page)
	fmt.Println(fifthPart)

	finalAnswer := firstPart * secondPart * thirdPart * fourthPart * fifthPart

	fmt.Println(finalAnswer)
}
