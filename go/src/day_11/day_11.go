package day11

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const PART_ONE_ITERATIONS = 25
const PART_TWO_ITERATIONS = 75

type Stones struct {
	cache map[[2]int]int
	// part_1 int
	// part_2 int
}

func DoDay11() (int, string, string) {
	part_1 := 0
	part_2 := 0

	file := helpers.GetFile(11)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stone_handler := Stones{make(map[[2]int]int)}

	for scanner.Scan() {
		for _, stone := range strings.Split(scanner.Text(), " ") {

			sone_val, err := strconv.Atoi(stone)
			if err != nil {
				fmt.Println(stone)
				log.Fatal("Should have been int")
			}
			part_2 += stone_handler.handleStone(sone_val, PART_TWO_ITERATIONS)
			part_1 += stone_handler.handleStone(sone_val, PART_ONE_ITERATIONS)

		}

	}

	return 11, strconv.Itoa(part_1), strconv.Itoa(part_2)
}

func (s *Stones) handleStone(stone_size int, iterations int) int {

	if iterations == 0 {
		return 1
	}

	key := [2]int{stone_size, iterations}

	if cached_values, cached := s.cache[key]; cached {
		return cached_values
	}

	next_iteration := iterations - 1

	var final_val int

	if stone_size == 0 {
		final_val = s.handleStone(1, next_iteration)
	} else {

		digit_len := len(strconv.Itoa(stone_size))

		if digit_len%2 == 0 {
			x := int(math.Pow10(digit_len / 2))

			final_val = +s.handleStone(stone_size/x, next_iteration) + s.handleStone(stone_size%x, next_iteration)
		} else {
			final_val = s.handleStone(stone_size*2024, next_iteration)
		}
	}

	s.cache[key] = final_val

	return final_val
}
