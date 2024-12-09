package day7

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"log"
	"math"
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

		if tryLine(target, values, false) {
			part_1 += target
		} else if tryLine(target, values, true) {
			part_2 += target
		}
	}

	part_2 += part_1

	return 7, strconv.Itoa(part_1), strconv.Itoa(part_2)
}

func tryLine(target int, values []int, allow_concat bool) bool {
	if len(values) == 1 {
		return target == values[0]
	}

	last_val := values[len(values)-1]

	if target%last_val == 0 && tryLine(target/last_val, values[:len(values)-1], allow_concat) {
		return true
	}

	if target-last_val >= 0 && tryLine(target-last_val, values[:len(values)-1], allow_concat) {
		return true
	}

	if allow_concat {

		digit_len := len(strconv.Itoa(last_val))
		ten_pow := int(math.Pow(float64(10), float64(digit_len)))

		if target%ten_pow == last_val && tryLine(target/ten_pow, values[:len(values)-1], true) {
			return true
		}
	}

	return false
}
