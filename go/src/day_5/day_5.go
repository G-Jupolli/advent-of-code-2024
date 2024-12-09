package day5

import (
	"advent_of_code_2024/helpers"
	"bufio"
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

		// Rules are always formatted as a|b where a and b are 2 digit ints
		// It's ok to black goose this and handle the failure if it ever happens
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

		// Rules are that int a must happen before b
		// The array is formatted with each index being 1/3 of int a so we need
		// to index into the arr at a * 3 and then the floor of int b / 32
		// This is done so we can bitwise check the b mod 32th bit
		lookup_idx := (left * 3) + (right / 32)
		lookup_bit := uint32(1 << (right % 32))

		// Update the bit here
		lookup[lookup_idx] = lookup[lookup_idx] | lookup_bit
	}

	for scanner.Scan() {
		line := scanner.Text()

		var pages []int

		// Manuals can have differnt length of pages so a split is ok here
		// There may be a quicker solution as the inputs are only ever
		// 2 digits long so skipping every 3rd rune could be quicker
		for _, page := range strings.Split(line, ",") {
			// Pre convert to int here
			page_int, err := strconv.Atoi(page)
			if err != nil {
				log.Fatal("Should be Int here")
			}

			pages = append(pages, page_int-10)
		}

		if is_ordered, middle_val := checkLine(lookup, pages); is_ordered {
			accumilator += middle_val
		} else {
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

			// The idea here is that if j must be before i then we can swap them.
			// In the recursion, if there ecxists n where j, n..., i
			// we check j against all n and swap where needed resulting
			// in the end being a complete array the follows all rules
			tmp := pages[i]
			pages[i] = pages[j]
			pages[j] = tmp

			_, p2 := checkLine(lookup, pages)

			return false, p2
		}
	}

	return true, pages[len(pages)/2] + 10
}
