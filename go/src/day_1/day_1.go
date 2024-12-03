package day1

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"log"
	"sort"
	"strconv"
	"strings"
)

func DoDay1() (int, string, string) {

	file := helpers.GetFile(1)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left_ids []int
	var right_ids []int
	right_id_map := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		string_parts := strings.Split(line, "   ")

		left, err := strconv.Atoi(string_parts[0])
		if err != nil {
			log.Fatal("Not int", err)
		}

		left_ids = append(left_ids, left)

		right, err := strconv.Atoi(string_parts[1])
		if err != nil {
			log.Fatal("Not int", err)
		}

		right_ids = append(right_ids, right)
		right_id_map[right]++
	}

	sort.Ints(left_ids)
	sort.Ints(right_ids)

	accumilator_first := 0
	accumilator_second := 0

	for i := 0; i < len(left_ids); i++ {
		left_id := left_ids[i]
		dif := left_id - right_ids[i]

		if dif < 0 {
			dif *= -1
		}

		accumilator_first += dif
		accumilator_second += left_id * right_id_map[left_id]
	}

	return 1, strconv.Itoa(accumilator_first), strconv.Itoa(accumilator_second)
}
