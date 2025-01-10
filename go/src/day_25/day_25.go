package day25

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"strconv"
)

/*

Schematic is for a lock if the first row is '#' e.g.
#####
.####		Can be written as heights
.####		0,5,3,4,3 by counting the '#'
.####
.#.#.
.#...
.....

And a key starts with '.' e.g.
.....
#....		Can be written as heights
#....		5,0,2,1,3 by counting the '#'
#...#
#.#.#
#.###
#####

Heights can only be 0 - 5 so they fit within 3 bits

This means that an item can be represented in a uint16
Where the first bit dictates if the item is a lock (1) or key (0)

Therefore the example lock can be: 	1_000_101_011_100_011
							key:	0_101_000_010_001_011

This means that to compare, we check if the last 3 bits added > 5
We then nee to right shifr them both by 3 and re compare
Until we have compared all the heights

We have to check each lock against each key so we save an array
of locks and an array of keys.
When we parse a key / lock, we check against the existing ones before inserting.
This way we're checking everything without needing to cross check everything at the end.

*/

const DAY = 25

const lock_flag = uint16(1 << 15)
const height_mask = item(0b0111)

const check_char = '#'

type item uint16

type cache struct {
	p1    int
	keys  []item
	locks []item
}

func DoDay() (int, string, string) {
	file := helpers.GetFile(DAY)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	accumilator_second := 0

	c := cache{
		0,
		[]item{},
		[]item{},
	}

	for {
		ok, is_lock, new_item := parse_item(scanner)

		if !ok {
			break
		}

		c.insert_item(is_lock, new_item)

	}

	return DAY, strconv.Itoa(c.p1), strconv.Itoa(accumilator_second)
}

func parse_item(scanner *bufio.Scanner) (bool, bool, item) {

	if !scanner.Scan() {
		return false, false, 0
	}

	is_lock := scanner.Bytes()[0] == check_char

	var vals [5]uint16

	for r := 0; r < 5; r += 1 {
		if !scanner.Scan() {
			panic("Line should exist")
		}

		line := scanner.Bytes()

		for i := 0; i < 5; i += 1 {
			if line[i] == check_char {
				vals[i] += 1
			}
		}
	}

	res := (vals[0] << 3) | vals[1]
	res = (res << 3) | vals[2]
	res = (res << 3) | vals[3]
	res = (res << 3) | vals[4]

	if is_lock {
		res |= lock_flag
	}

	if !scanner.Scan() {
		panic("Line should exist pa")
	}
	scanner.Scan()

	return true, is_lock, item(res)
}

func compare_vals(key item, lock item) bool {

	// We need to check all 5 3 bit chunks
	for i := 0; i < 5; i += 1 {
		key_val := (key >> (i * 3)) & height_mask
		lock_val := (lock >> (i * 3)) & height_mask

		if key_val+lock_val > 5 {
			return false
		}
	}

	return true
}

func (c *cache) insert_item(is_lock bool, val item) {
	if is_lock {
		c.locks = append(c.locks, val)

		for _, key := range c.keys {
			if compare_vals(key, val) {
				c.p1 += 1
			}
		}
	} else {
		c.keys = append(c.keys, val)

		for _, lock := range c.locks {
			if compare_vals(val, lock) {
				c.p1 += 1
			}
		}
	}
}
