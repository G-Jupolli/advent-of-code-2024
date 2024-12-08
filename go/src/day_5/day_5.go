package day5

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

Building an efficient lookup array.

The data appears to all be integers between 10 -> 99.
This means we have a possible 90 numbers.
These would fit within the bits of 3 u32.

This means that we can initiate a 89 * 3 = 270 len u32 aray.
This gives us a storage size of 8544bits.

By doing this, we have no need to hold onto and strings or any hashmaps
We then easily index into the array for any lookup we need.
*/

func DoDay5() (int, string, string) {

	file := helpers.GetFile(5)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lookup [269]uint32

	accumilator := 0
	accumilator_2 := 0

	// Handle the ordering
	for scanner.Scan() {
		line := scanner.Text()

		// This is where we switch to the pages
		if len(line) < 4 {
			break
		}

		left, err := strconv.Atoi(line[:2])
		if err != nil {
			log.Fatal("Should be Int here")
		}

		right, err := strconv.Atoi(line[3:5])
		if err != nil {
			log.Fatal("Should be Int here")
		}

		left -= 10
		right -= 10

		lookup_idx := (left * 3) + (right / 32)
		lookup_bit := uint32(1 << (right % 32))

		lookup[lookup_idx] = lookup[lookup_idx] | lookup_bit
	}

	for scanner.Scan() {
		line := scanner.Text()

		var pages []int

		for _, page := range strings.Split(line, ",") {
			page_int, err := strconv.Atoi(page)
			if err != nil {
				log.Fatal("Should be Int here")
			}

			pages = append(pages, page_int-10)
		}

		if is_ordered, middle_val := checkLine(lookup, pages); is_ordered {
			accumilator += middle_val
		} else {
			fmt.Println("l2", middle_val)
			accumilator_2 += middle_val
		}
	}

	return 5, strconv.Itoa(accumilator), strconv.Itoa(accumilator_2)
}

func checkLine(lookup [269]uint32, pages []int) (bool, int) {

	// No need to check last element in i
	// Part1_Loop:
	for i := 0; i < (len(pages) - 1); i++ {
		lookup_idx := pages[i] * 3

		for j := i + 1; j < len(pages); j++ {
			right_addr := pages[j] / 32
			lookup_bit := uint32(1 << (pages[j] % 32))

			// Check if i -> j is ok
			if (lookup[lookup_idx+right_addr] & lookup_bit) != 0 {
				continue
			}

			// fmt.Println("Check 2 ")

			new_lookup_idx := pages[j] * 3
			new_right_addr := pages[i] / 32
			new_lookup_bit := uint32(1 << (pages[i] % 32))

			// Check if j -> i is ok
			if (lookup[new_lookup_idx+new_right_addr] & new_lookup_bit) == 0 {
				// If not then there is no safe config
				return false, 0
			}

			tmp := pages[i]
			pages[i] = pages[j]
			pages[j] = tmp

			return checkLine(lookup, pages)
		}
	}

	return true, pages[len(pages)/2] + 10
}
