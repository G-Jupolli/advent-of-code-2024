package day2

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"log"
	"strconv"
	"strings"
)

func DoDay2() (int, string, string) {
	file := helpers.GetFile(2, true)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	first_safe_reactors := 0
	second_safe_reactors := 0

	for scanner.Scan() {
		line := scanner.Text()

		reactor_parts := strings.Split(line, " ")

		var reactor_part_ints []int

		for _, part := range reactor_parts {
			part_int, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			reactor_part_ints = append(reactor_part_ints, part_int)
		}

		is_safe, fail_idx := isSafe(reactor_part_ints)

		if is_safe {
			first_safe_reactors++
			second_safe_reactors++
			// Special case where the 2nd item fails
		} else if fail_idx == 1 && isAcceptable(reactor_part_ints, 0) {

			second_safe_reactors++

		} else if isAcceptable(reactor_part_ints, fail_idx) || isAcceptable(reactor_part_ints, fail_idx+1) {

			second_safe_reactors++
		}
	}

	return 1, strconv.Itoa(first_safe_reactors), strconv.Itoa(second_safe_reactors)
}

func isSafe(parts []int) (bool, int) {
	// Probably never happens
	if len(parts) < 2 {
		return true, 0
	}

	is_increasing := true

	for i := 0; i < len(parts)-1; i++ {
		left := parts[i]
		right := parts[i+1]

		dif := left - right
		did_increase := dif < 0

		if dif < 0 {
			dif *= -1
		}

		if dif < 1 || dif > 3 {
			return false, i
		}

		if i == 0 {
			is_increasing = did_increase
		} else if is_increasing != did_increase {
			return false, i
		}

	}

	return true, 0
}

func isAcceptable(parts []int, fail_idx int) bool {
	// Failing on the last reading is acceptable
	if fail_idx == len(parts)-1 {
		return true
	}

	var new_reactor_parts []int

	for i, p := range parts {
		if i != fail_idx {
			new_reactor_parts = append(new_reactor_parts, p)
		}
	}

	is_safe, _ := isSafe(new_reactor_parts)

	return is_safe
}
