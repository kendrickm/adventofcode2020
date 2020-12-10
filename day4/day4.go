package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var filename = "input.dat"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type inputBlock struct {
	inputs []string
}

type field struct {
	key   string
	value string
}

type parsedBlock struct {
	fields []field
	keys   []string
}

func parseInputFile() []*inputBlock {

	var blocks []*inputBlock
	dat, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(dat), "\n")
	block := &inputBlock{}
	for _, line := range lines {
		if line == "" {
			// fmt.Println("Blank line")
			if block == nil {
				panic("Block is nil somehow")
			}
			blocks = append(blocks, block)
			block = &inputBlock{}
		} else {
			block.inputs = append(block.inputs, line)
		}

	}
	return blocks
}

func parseInputBlocks(input []string) *parsedBlock {
	var values []string

	for _, item := range input {
		x := strings.Split(item, " ")
		for _, y := range x {
			values = append(values, y)
		}
	}

	parsed := &parsedBlock{}
	for _, item := range values {
		colonSplit := strings.Split(item, ":")
		f := field{colonSplit[0], colonSplit[1]}
		parsed.fields = append(parsed.fields, f)
		parsed.keys = append(parsed.keys, colonSplit[0])
	}

	return parsed

}

func (block *parsedBlock) validateEntry() bool {
	validators := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		"cid"}
	errors := 0
	for _, val := range validators {
		found := Contains(block.keys, val)

		if !found {
			// fmt.Println("Error")
			if val == "cid" {
				// fmt.Println("We can skip cid")
			} else {
				errors++
			}
		} else {
			fmt.Println(val)
			switch val {
			case "byr":
				if !findField(block.fields, "byr").validateBYR() {
					errors++
				}
			case "iyr":
				if !findField(block.fields, "iyr").validateIYR() {
					errors++
				}
			case "eyr":
				if !findField(block.fields, "eyr").validateEYR() {
					errors++
				}
			case "hgt":
				if !findField(block.fields, "hgt").validateHGT() {
					errors++
				}
			case "hcl":
				if !findField(block.fields, "hcl").validateHCL() {
					errors++
				}
			case "ecl":
				if !findField(block.fields, "ecl").validateECL() {
					errors++
				}
			case "pid":
				if !findField(block.fields, "pid").validatePID() {
					errors++
				}
			case "cid":
			}
		}
	}

	return errors == 0

}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func findField(a []field, x string) field {
	n := field{}
	for _, n = range a {
		if x == n.key {
			return n
		}
	}
	return n
}

func (f field) validateBYR() bool {
	valid := true
	value, err := strconv.Atoi(f.value)
	check(err)
	if value < 1920 || value > 2002 {
		valid = false
	}

	return valid
}

func (f field) validateIYR() bool {
	valid := true
	value, err := strconv.Atoi(f.value)
	check(err)
	if value < 2010 || value > 2020 {
		valid = false
	}

	return valid
}

func (f field) validateEYR() bool {
	valid := true
	value, err := strconv.Atoi(f.value)
	check(err)
	if value < 2020 || value > 2030 {
		valid = false
	}

	return valid
}

func (f field) validateHGT() bool {
	valid := false
	rNumIn, _ := regexp.Compile("^[0-9][0-9]in")
	rNumCm, _ := regexp.Compile("^[0-9][0-9][0-9]cm")
	rIn, _ := regexp.Compile("in")
	rCm, _ := regexp.Compile("cm")

	if rNumIn.MatchString(f.value) {
		height, err := strconv.Atoi(rIn.ReplaceAllString(f.value, ""))
		check(err)
		valid = (height <= 76 && height >= 59)
	} else if rNumCm.MatchString(f.value) {
		height, err := strconv.Atoi(rCm.ReplaceAllString(f.value, ""))
		check(err)
		valid = (height <= 193 && height >= 150)
	}
	return valid
}

func (f field) validateHCL() bool {
	rHCL, _ := regexp.Compile("^#[0-9a-f]{6}")
	valid := rHCL.MatchString(f.value)

	return valid
}

func (f field) validateECL() bool {
	valid := false
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, c := range colors {
		if f.value == c {
			valid = true
			break
		}
	}
	return valid
}

func (f field) validatePID() bool {

	valid := len(f.value) == 9

	return (valid)

}

func main() {

	blocks := parseInputFile()
	errors := 0

	parsedFile := []*parsedBlock{}
	for _, x := range blocks {
		value := parseInputBlocks(x.inputs)
		parsedFile = append(parsedFile, value)
		if !value.validateEntry() {
			errors++
		}
	}
	length := len(blocks)
	correct := length - errors
	fmt.Printf("Total number of entries %d \n", length)
	fmt.Printf("Total number of errors %d \n", errors)
	fmt.Printf("Total number of correct %d \n", correct)
}
