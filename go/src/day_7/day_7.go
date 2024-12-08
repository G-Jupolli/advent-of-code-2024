package day7

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
Strategy
*/

func DoDay7() (int, string, string) {

	file := helpers.GetFile(7)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part_1 := 0
	part_2 := -1

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

		line_val := checkLine(target, values)

		if line_val == 0 {
			fmt.Println(line)
			// fmt.Println("T: ", target, ", Vals : ", values)
		}

		part_1 += line_val
	}

	return 7, strconv.Itoa(part_1), strconv.Itoa(part_2)
}

func checkLine(target int, values []int) int {
	// if len(values) < 1 {
	// 	return 0
	// }

	// The different combinations of + & * there can be
	variations := (1 << (len(values) - 1))
	// Start at 1 instead of 2 here as above validates if
	// target is the sum of the values
	checker := 0

	// is_valid := false
	// 14711933466277
	// 14652671737047
	// 14711933466277
	for checker < variations {

		accumilator := values[0]
		// fmt.Print(accumilator)

		for i := 1; i < len(values); i += 1 {
			mul_flag := 1 << (i - 1)

			if (mul_flag & checker) != 0 {
				// fmt.Print(" * ", values[i])
				accumilator *= values[i]
			} else {
				// fmt.Print(" + ", values[i])
				accumilator += values[i]
			}
		}

		// fmt.Print(" = ", accumilator)

		// fmt.Println()

		if accumilator == target {
			return target
		}

		checker += 1
	}

	fmt.Println("Tries ", checker)
	// 81 + 40 * 27 and 81 * 40 + 27 is 3267
	// println("len flags", variations)

	return 0
}
