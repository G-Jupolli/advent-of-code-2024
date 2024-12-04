package day4

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"strconv"
)

/*
Strategy

Finding "XMAS" in a word search.

The way I plan to do this is to have a size 4 array of arrays
These will each hold a line of the input
On 'X' or 'S' we seach Down Left, Down, Down Right, Right
By looking for these runes we are able to account for all words in all formats

p - [ Line 1
	  Line 2
	  Line 3
	  Line 4
    ]
Where P is a pointer to the index the line is on
Scan line at p an check words; if trying to go down and line is nil, load line in
Once we scan the line, line at index p is updated to next line and incr p
Therefore Line 1 -> Line 5 and p++

In doing this we need to access the array with ( p + x ) % 4
We also limit our memory use by only having 4 lines loaded in at any given time

*/

// var chars = [4]rune{'X', 'M', 'A', 'S'}
var incr_runes = [3]rune{'M', 'A', 'S'}
var decr_runes = [3]rune{'A', 'M', 'X'}

func DoDay4() (int, string, string) {

	file := helpers.GetFile(4)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialise the first 4 lines
	// Mkaing the assumption we have min 4 lines is ok here ( for now )
	container := [4]*[]rune{
		getNextLine(scanner),
		getNextLine(scanner),
		getNextLine(scanner),
		getNextLine(scanner),
	}

	// This pointer should be max 3 and ring buf around container
	line_idx := 0

	// Set this on getNextLine -> nil so that we don't need to bother checking down
	did_end := false
	did_end_p2 := false

	accumilator := 0
	accumilator_p2 := 0

	for {
		// This happens when all lines are processed
		// This allows the array to become all nil to define exit
		if container[line_idx] == nil {
			break
		}

		line_data := *container[line_idx]
		// line_data := *line_data

		container[line_idx] = nil

		// fmt.Println("line", line_data)
		// fmt.Println("Should be nil", container[line_idx])

		for i, itm := range line_data {
			switch itm {
			case 'X':
				accumilator += checkPart1(container, line_data, line_idx, i, did_end, true)
			case 'S':
				accumilator += checkPart1(container, line_data, line_idx, i, did_end, false)
				if checkPArt2(container, line_data, itm, line_idx, i, did_end_p2) {
					accumilator_p2++
				}
			case 'M':
				if checkPArt2(container, line_data, itm, line_idx, i, did_end_p2) {
					accumilator_p2++
				}
			default:
				continue
			}

		}

		if !did_end {
			new_line_data := getNextLine(scanner)

			if new_line_data == nil {
				did_end = true
			} else {
				container[line_idx] = new_line_data
			}
		} else {
			// p2 can have 1 more line below p1
			did_end_p2 = true
		}

		line_idx = (line_idx + 1) % 4
	}

	return 4, strconv.Itoa(accumilator), strconv.Itoa(accumilator_p2)
}

func getNextLine(scanner *bufio.Scanner) *[]rune {
	has_more := scanner.Scan()

	if !has_more {
		return nil
	}

	line := []rune(scanner.Text())

	return &line
}

func checkPart1(container [4]*[]rune, line []rune, line_idx int, char_idx int, did_end bool, is_start bool) int {
	matches := 0

	var check_runes [3]rune

	if is_start {
		check_runes = incr_runes
	} else {
		check_runes = decr_runes
	}

	// Can remove right checks if there are not enough chars
	check_right := len(line)-char_idx >= 4

	// Check Right
	if check_right &&
		line[char_idx+1] == check_runes[0] &&
		line[char_idx+2] == check_runes[1] &&
		line[char_idx+3] == check_runes[2] {
		matches++
	}

	// In this case there are < 3 lines below this point
	// So there cannot be any downward matches
	if did_end {
		return matches
	}

	first := *container[(line_idx+1)%4]
	second := *container[(line_idx+2)%4]
	third := *container[(line_idx+3)%4]

	// Check Down
	if first[char_idx] == check_runes[0] &&
		second[char_idx] == check_runes[1] &&
		third[char_idx] == check_runes[2] {
		matches++
	}

	// Check Down Righ
	if check_right &&
		first[char_idx+1] == check_runes[0] &&
		second[char_idx+2] == check_runes[1] &&
		third[char_idx+3] == check_runes[2] {
		matches++
	}

	if char_idx < 3 {
		return matches
	}

	// Check Down Left
	if first[char_idx-1] == check_runes[0] &&
		second[char_idx-2] == check_runes[1] &&
		third[char_idx-3] == check_runes[2] {
		matches++
	}

	return matches
}

// Only try to check from top left so boll return is fine here
func checkPArt2(container [4]*[]rune, line []rune, char rune, line_idx int, char_idx int, did_end bool) bool {
	if did_end || (len(line)-char_idx) <= 2 {
		return false
	}

	//   a _ b
	//   _ c _
	//   d _ e

	// where a is char

	top_char := line[char_idx+2]

	if top_char != 'M' && top_char != 'S' {
		return false
	}

	// Case when
	//   M _ M
	is_top := top_char == char

	middle_line := *container[(line_idx+1)%4]
	if middle_line[char_idx+1] != 'A' {
		return false
	}

	last_line := *container[(line_idx+2)%4]

	bottom_char := last_line[char_idx]

	if bottom_char != 'M' && bottom_char != 'S' {
		return false
	}

	// Handles case
	//   M _ M
	//   _ _ _
	//   M _ _
	if (bottom_char == char) == is_top {
		return false
	}

	right_char := last_line[char_idx+2]

	if right_char != 'M' && right_char != 'S' {
		return false
	}

	// Handles case
	//   M _ _
	//   _ _ _
	//   M _ M
	if right_char == bottom_char && !is_top {
		return false
	}

	// Handles case
	//   M _ M
	//   _ _ _
	//   _ _ M
	if (right_char == top_char) == is_top {
		return false
	}

	return true
}
