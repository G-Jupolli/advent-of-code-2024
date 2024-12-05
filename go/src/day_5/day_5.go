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

		fmt.Printf("LB - %032b\n", lookup_bit)
		println("Left - ", left, ", Right - ", right, ", Lookup idx - ", lookup_idx, ", From lookup bit", lookup_bit)

		lookup[lookup_idx] = lookup[lookup_idx] | lookup_bit

		fmt.Printf("%v\n", lookup)

	}

	for scanner.Scan() {
		line := scanner.Text()

		if is_ordered, middle_val := checkLine(lookup, line); is_ordered {
			fmt.Println("MV", middle_val)
			accumilator += middle_val
		}
	}

	return 5, strconv.Itoa(accumilator), "boop"
}

func checkLine(lookup [269]uint32, line string) (bool, int) {

	var pages []int

	for _, page := range strings.Split(line, ",") {
		page_int, err := strconv.Atoi(page)
		if err != nil {
			log.Fatal("Should be Int here")
		}

		pages = append(pages, page_int-10)
	}

	// No need to check last element in i
	for i := 0; i < (len(pages) - 1); i++ {
		lookup_idx := pages[i] * 3

		for j := i + 1; j < len(pages); j++ {
			right_addr := pages[j] / 32
			lookup_bit := uint32(1 << (pages[j] % 32))

			if lookup[lookup_idx+right_addr]&lookup_bit == 0 {
				return false, 0
			}
		}
	}

	// Need to give back the 10 here
	return true, pages[len(pages)/2] + 10
}
