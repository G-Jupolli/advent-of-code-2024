package day7

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"log"
	"strconv"
	"strings"
)

/*
Strategy

Mapping over all of the possible methods is crazy inefficient.
It would be better to backtrack through the values as if the target is a
multiple of a value, it is fine to consider it a multiple unless proven wrong later
*/

func DoDay7() (int, string, string) {

	file := helpers.GetFile(7)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part_1 := 0
	part_2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		line_data := strings.Split(line, ": ")

		target, err := strconv.Atoi(line_data[0])
		if err != nil {
			log.Fatal("Should be Int here")
		}

		var values []int

		for _, v := range strings.Split(line_data[1], " ") {
			v_int, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Should be Int here")
			}
			values = append(values, v_int)
		}

		if tryLine(target, values) {
			part_1 += target
		}

		// if p1_ok, line_val := checkLine(target, values); p1_ok {
		// 	part_1 += line_val
		// } else {
		// 	part_2 += line_val
		// }
	}

	part_2 += part_1

	return 7, strconv.Itoa(part_1), strconv.Itoa(part_2)
}

func checkLine(target int, values []int) (bool, int) {

	// The different combinations of + & * there can be
	variations := (1 << (len(values) - 1))
	// Start at 1 instead of 2 here as above validates if
	// target is the sum of the values
	checker := 0

	for checker < variations {

		accumilator := values[0]

		for i := 1; i < len(values); i += 1 {
			mul_flag := 1 << (i - 1)

			if (mul_flag & checker) != 0 {
				accumilator *= values[i]
			} else {
				accumilator += values[i]
			}
		}

		if accumilator == target {
			return true, target
		}

		checker += 1
	}

	return false, checkPart2(target, values)
	// return false, 0
}

func checkPart2(target int, values []int) int {

	// The different combinations of + & * there can be
	variations := (1 << (len(values) - 1))
	// Start at 1 instead of 2 here as above validates if
	// target is the sum of the values
	checker := 0

	for checker < variations {
		for x := 1; x < variations; x += 1 {

			accumilator := values[0]

			// fmt.Print(accumilator)

			for i := 1; i < len(values); i += 1 {
				curr_val := values[i]
				flag_pos := 1 << (i - 1)

				if (flag_pos & x) != 0 {
					// print(" || ", curr_val)
					accumilator = (accumilator * (len(strconv.Itoa(curr_val)) * 10)) + curr_val
				} else if (flag_pos & checker) != 0 {
					// print(" * ", curr_val)
					accumilator *= curr_val
				} else {
					// print(" + ", curr_val)
					accumilator += curr_val
				}
			}

			// fmt.Print(" = ", accumilator)
			// fmt.Println()

			if accumilator == target {
				return target
			}

		}
		checker += 1
	}

	return 0
}

func tryLine(target int, values []int) bool {
	if len(values) == 1 {
		return target == values[0]
	}

	last_val := values[len(values)-1]

	if target%last_val == 0 && tryLine(target/last_val, values[:len(values)-1]) {
		return true
	}

	if target-last_val >= 0 && tryLine(target-last_val, values[:len(values)-1]) {
		return true
	}

	return false
}
