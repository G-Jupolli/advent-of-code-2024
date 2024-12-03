package day3

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"log"
	"regexp"
	"strconv"
)

// Three way match group for mul([number], [numnber]), do() and don't()
// The mul capture also has an additional capture group for the nummbers for minor convenience
const regex_pattern = `(?m)(mul\((\d+),(\d+)\)|do\(\)|don\'t\(\))`

func DoDay3() (int, string, string) {

	var re = regexp.MustCompile(regex_pattern)
	file := helpers.GetFile(3, true)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	accumilator_first := 0
	accumilator_second := 0

	is_do := true

	for scanner.Scan() {
		line := scanner.Text()
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			// First 3 letters is all we need to differentiate the functions
			function_match := match[0][:3]

			switch function_match {
			case "mul":
				left, err := strconv.Atoi(match[2])
				if err != nil {
					log.Fatal(err)
				}
				right, err := strconv.Atoi(match[3])
				if err != nil {
					log.Fatal(err)
				}

				addition := left * right
				// Part 1 is always accumulate
				accumilator_first += addition

				if is_do {
					accumilator_second += addition
				}
			case "do(":
				is_do = true
			case "don":
				is_do = false
			default:
				log.Fatal("Invalid regex match", match)
			}
		}
	}

	return 3, strconv.Itoa(accumilator_first), strconv.Itoa(accumilator_second)
}
