package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	min    int
	max    int
	letter string
}

type entry struct {
	policy   *passwordPolicy
	password string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput() []*entry {
	dat, err := ioutil.ReadFile("day2.input")
	check(err)
	input := strings.Split(string(dat), "\n")

	entries := make([]*entry, len(input)-1)
	//14-15 h: hhhhhhhhhhhhhjh
	for i := range entries {
		s := input[i]
		pp := &passwordPolicy{}
		segments := strings.Split(s, " ")
		if segments[0] == "" {
			break
		}
		pp.min, err = strconv.Atoi(strings.Split(segments[0], "-")[0])
		check(err)
		pp.max, err = strconv.Atoi(strings.Split(segments[0], "-")[1])
		check(err)
		pp.letter = strings.Trim(segments[1], ":")
		password := segments[2]
		e := &entry{pp, password}
		entries[i] = e
	}
	return entries
}

func (e *entry) validatePolicyPart1() bool {
	characterCount := strings.Count(e.password, e.policy.letter)
	// fmt.Println(characterCount)
	return ((characterCount >= e.policy.min) && (characterCount <= e.policy.max))
}

func (e *entry) validatePolicyPart2() bool {
	firstNumToCheck := e.policy.min - 1
	secondNumToCheck := e.policy.max - 1
	firstChar := string(e.password[firstNumToCheck])
	secondChar := string(e.password[secondNumToCheck])

	// fmt.Printf("Character at space %d ", firstNumToCheck)
	// fmt.Println(firstChar)
	//
	// fmt.Printf("Character at space %d ", secondNumToCheck)
	// fmt.Println(secondChar)

	results := false

	if firstChar == e.policy.letter {
		if secondChar != e.policy.letter {
			results = true
		}
	}

	if secondChar == e.policy.letter {
		if firstChar != e.policy.letter {
			results = true
		}
	}

	return results
}

func main() {
	entries := parseInput()
	numCorrect := 0
	for _, e := range entries {
		if e.validatePolicyPart1() {
			numCorrect++
		}
	}
	//
	fmt.Println(numCorrect)

	numCorrect2 := 0
	for _, e := range entries {
		if e.validatePolicyPart2() {
			numCorrect2++
		}
	}

	fmt.Println(numCorrect2)
	// fmt.Println("Validating the parsing logic")
	// fmt.Println(entries[24].policy.min)
	// fmt.Println(entries[24].policy.max)
	// fmt.Println(entries[24].policy.letter)
	// fmt.Println(entries[24].password)
	// entries[24].validatePolicyPart2()
}
