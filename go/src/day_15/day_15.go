package day15

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"fmt"
	"log"
	"strconv"
)

const ROBOT = '@'
const BOX = 'O'
const WALL = '#'

const MOVE_UP = '^'
const MOVE_DOWN = 'v'
const MOVE_LEFT = '<'
const MOVE_RIGHT = '>'

type Point struct {
	x int
	y int
}

/*
The memory savings here are probably neglidgable in scope due to the ability
of modern ram, I just want to do it this way to practice bit manipulation.
*/
type Arena struct {
	width int
	depth int
	robot Point
	boxes []uint64
	walls []uint64
}

func DoDay15() (int, string, string) {
	file := helpers.GetFile(15)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	accumilator_first := 0
	accumilator_second := 0

	arena := parseArena(scanner)

	for scanner.Scan() {
		line := scanner.Text()

		for _, command := range line {
			// fmt.Println(command)

			arena.handleCommand(command)
		}
	}

	for i, box_row := range arena.boxes {
		// If no boxes exist on row can just skip
		if box_row == 0 {
			continue
		}

		for j := 0; j < arena.width-1; j += 1 {
			flag := uint64(1 << j)

			// Check if curr pos has a box
			if box_row&flag == 0 {
				continue
			}
			accumilator_first += (100 * (i + 1)) + j + 1
		}
	}
	fmt.Println()
	return 15, strconv.Itoa(accumilator_first), strconv.Itoa(accumilator_second)
}

func parseArena(scanner *bufio.Scanner) Arena {
	// Do first scan to get past initial wall
	scanner.Scan()
	// Load fist line we care about in
	scanner.Scan()

	prev := scanner.Text()

	var robot Point
	var boxes []uint64
	var walls []uint64

	var width int
	depth := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Checks for the partition to know when the arena is all loaded
		if len(line) < 5 {
			break
		}

		boxes = append(boxes, 0)
		walls = append(walls, 0)

		// We dont care about the arena walls here
		// We parese prev so we dont parse in the Areana upper and lower bounds
		for i := 1; i < len(prev)-1; i += 1 {
			switch prev[i] {
			case ROBOT:
				robot = Point{i - 1, depth}
			case WALL:
				bit_flag := uint64(1 << (i - 1))
				walls[depth] |= bit_flag
			case BOX:
				bit_flag := uint64(1 << (i - 1))
				boxes[depth] |= bit_flag
			}
		}

		depth += 1
		prev = line
		// Not including the arena walls
		width = len(line) - 1
	}

	return Arena{
		width,
		depth,
		robot,
		boxes,
		walls,
	}
}

func (a *Arena) handleCommand(command rune) {

	delta_x := 0
	delta_y := 0

	switch command {

	case MOVE_RIGHT:
		delta_x = 1
	case MOVE_LEFT:
		delta_x = -1
	case MOVE_UP:
		delta_y = -1
	case MOVE_DOWN:
		delta_y = 1

	default:
		log.Fatal("invalid command ", command)
	}

	can_move, empty_slot := a.find_empty_slot(delta_x, delta_y)

	if !can_move {
		return
	}

	// Robot position does not affect boxes so can pre move here
	a.robot = Point{
		a.robot.x + delta_x,
		a.robot.y + delta_y,
	}

	if empty_slot.x == a.robot.x && empty_slot.y == a.robot.y {
		return
	}

	switch command {
	case MOVE_RIGHT:
		fallthrough
	case MOVE_LEFT:
		// This makes a 1 at pos where first object moves and last object moves into
		object_changes := uint64((1 << a.robot.x) | (1 << empty_slot.x))

		// Use a XOR here to turn off the first and the empty space on
		a.boxes[a.robot.y] ^= object_changes

	case MOVE_UP:
		fallthrough
	case MOVE_DOWN:
		// Make a flag repr for the column to check object on
		col_flag := uint64(1 << a.robot.x)

		// Use XOR here to turn object above robot off and object in empty space on
		a.boxes[a.robot.y] ^= col_flag
		a.boxes[empty_slot.y] ^= col_flag
	}
}

func (a *Arena) find_empty_slot(delta_x int, delta_y int) (bool, Point) {
	check_x := a.robot.x + delta_x
	check_y := a.robot.y + delta_y

	for {
		if check_x < 0 || check_y < 0 || check_x == a.width-1 || check_y == a.depth {
			return false, Point{}
		}

		is_wall, is_object := a.check_pos(check_x, check_y)

		if is_wall {
			return false, Point{}
		}

		if !is_object {
			return true, Point{
				check_x,
				check_y,
			}
		}

		check_x += delta_x
		check_y += delta_y
	}
}

func (a *Arena) check_pos(x int, y int) (bool, bool) {
	column_flag := uint64(1 << x)

	is_wall := a.walls[y]&column_flag != 0

	if is_wall {
		return is_wall, false
	}

	is_object := a.boxes[y]&column_flag != 0

	return false, is_object

}
