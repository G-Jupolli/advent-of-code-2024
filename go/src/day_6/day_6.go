package day6

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"strconv"
)

/*
Strategy

Part 2
Seems like if I look at the direction 90 deg from where I currently am
Basically emulating an obstruction
*/

func DoDay6() (int, string, string) {

	file := helpers.GetFile(6)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var board_state [][3]uint64
	var visited_positions [][3]uint64
	var board_len int

	var guard_x int
	var guard_y int

	// 0 -> North
	// 1 -> East
	// 2 -> South
	// 3 -> West
	var guard_heading int32 = 0

	for scanner.Scan() {
		line := scanner.Text()

		board_row, guard_pos := parseLine(line)
		board_state = append(board_state, board_row)
		// Keep visited positions with same len as board state
		var new_unvisited [3]uint64
		visited_positions = append(visited_positions, new_unvisited)

		if guard_pos > -1 {
			board_len = len(line)
			guard_x = guard_pos
			guard_y = len(board_state) - 1
			// break
		}
	}

	visited_count := 0

	for {
		x_pos_flag := uint64(1 << (guard_x % 64))

		is_visited := visited_positions[guard_y][guard_x/64]&x_pos_flag != 0

		new_x := guard_x
		new_y := guard_y

		switch guard_heading {
		case 0:
			new_y -= 1
		case 1:
			new_x += 1
		case 2:
			new_y += 1
		case 3:
			new_x -= 1
		}

		if !checkIsInside(new_x, new_y, board_len-1, len(board_state)-1) {
			if !is_visited {
				visited_count += 1
			}
			break
		}

		is_obstructed := checkObstructed(board_state, new_x, new_y)

		if is_obstructed {
			guard_heading = (guard_heading + 1) % 4
		} else {
			if !is_visited {
				visited_positions[guard_y][guard_x/64] |= x_pos_flag
				visited_count += 1
			}

			guard_x = new_x
			guard_y = new_y

		}
	}

	accumilator_2 := -1

	return 6, strconv.Itoa(visited_count), strconv.Itoa(accumilator_2)
}

func parseLine(line string) ([3]uint64, int) {
	guard_idx := -1
	var flags [3]uint64

	for idx, char := range line {
		switch char {
		case '#':
			flags[idx/64] |= uint64(1 << (idx % 64))
		case '^':
			guard_idx = idx
		}
	}

	return flags, guard_idx
}

func checkObstructed(board_state [][3]uint64, x int, y int) bool {
	x_pos_flag := uint64(1 << (x % 64))
	board_row := board_state[y]
	item := board_row[x/64]

	return item&x_pos_flag != 0
}

func checkIsInside(x int, y int, max_x int, max_y int) bool {
	return x >= 0 && y >= 0 && x <= max_x && y <= max_y
}
