package day6

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"log"
	"strconv"
)

/*
Strategy
*/

func DoDay6() (int, string, string) {

	file := helpers.GetFile(6)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var board_state [][3]uint64
	var visited_positions [][3]uint64
	var board_len int

	var guard_x int
	var guard_y int = -1

	// 0 -> North
	// 1 -> East
	// 2 -> South
	// 3 -> West
	var guard_heading int32 = 0

	for scanner.Scan() {
		guard_y += 1
		line := scanner.Text()

		board_row, guard_pos := parseLine(line)
		board_state = append(board_state, board_row)
		// Keep visited positions with same len as bnoard state
		var new_unvisited [3]uint64
		visited_positions = append(visited_positions, new_unvisited)

		// No need to scan more once we have the guard position
		if guard_pos > -1 {
			board_len = len(line)
			guard_x = guard_pos
			break
		}
	}

	visited_count := 0

main_loop:
	for {
		x_pos_flag := uint64(1 << (guard_x % 64))

		is_visited := visited_positions[guard_y][guard_x/64]&x_pos_flag != 0

		switch guard_heading {
		// North
		case 0:
			// Top of area
			if guard_y == 0 {
				if !is_visited {
					visited_count += 1
				}
				break main_loop
			}

			// Check if there is obstruction above
			if board_state[guard_y-1][guard_x/64]&x_pos_flag != 0 {
				guard_heading = 1
				continue main_loop
			}

			// When moving off tile, check visited state
			if !is_visited {
				visited_positions[guard_y][guard_x/64] |= x_pos_flag
				visited_count += 1
			}

			guard_y -= 1

		// East
		case 1:
			// Going Off Edge
			if guard_x == board_len {
				if !is_visited {
					visited_count += 1
				}
				break main_loop
			}

			new_x_pos_flag := uint64(1 << ((guard_x % 64) + 1))

			// Check obstruction
			if board_state[guard_y][(guard_x+1)/64]&new_x_pos_flag != 0 {
				guard_heading = 2
				continue main_loop
			}

			// When moving off tile, check visited state
			if !is_visited {
				visited_positions[guard_y][guard_x/64] |= x_pos_flag
				visited_count += 1
			}

			guard_x += 1

		// West
		case 3:
			// Going Off Edge
			if guard_x == 0 {
				if !is_visited {
					visited_count += 1
				}
				break main_loop
			}

			new_x_pos_flag := uint64(1 << ((guard_x - 1) % 64))

			// Check obstruction
			if board_state[guard_y][(guard_x-1)/64]&new_x_pos_flag != 0 {
				guard_heading = 0
				continue main_loop
			}

			// When moving off tile, check visited state
			if !is_visited {
				visited_positions[guard_y][guard_x/64] |= x_pos_flag
				visited_count += 1
			}

			guard_x -= 1

		// South
		case 2:
			// Need to load in more board in this case
			if guard_y == len(board_state)-1 {

				// When there is no more, we move off board
				if !scanner.Scan() {
					if !is_visited {
						visited_count += 1
					}
					break main_loop
				}

				line := scanner.Text()

				board_row, _ := parseLine(line)
				board_state = append(board_state, board_row)
				// Keep visited positions with same len as bnoard state
				var new_unvisited [3]uint64
				visited_positions = append(visited_positions, new_unvisited)
			}

			// Check obstruction
			if board_state[guard_y+1][guard_x/64]&x_pos_flag != 0 {
				guard_heading = 3
				continue main_loop
			}

			// When moving off tile, check visited state
			if !is_visited {
				visited_positions[guard_y][guard_x/64] |= x_pos_flag
				visited_count += 1
			}

			guard_y += 1

		default:
			log.Fatal("Guard Heading Invalid")
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
