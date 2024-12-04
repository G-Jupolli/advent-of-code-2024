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

	accumilator := 0

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
			// println("Idx : ", i, ", Char : ", itm)

			var check_runes [3]rune

			switch itm {
			case 'X':
				check_runes = incr_runes
			case 'S':
				check_runes = decr_runes
			default:
				continue
			}

			// Can remove right checks if there are not enough chars
			check_right := len(line_data)-i >= 4

			// Check Right
			if check_right &&
				line_data[i+1] == check_runes[0] &&
				line_data[i+2] == check_runes[1] &&
				line_data[i+3] == check_runes[2] {
				accumilator++
			}

			// In this case there are < 3 lines below this point
			// So there cannot be any downward matches
			if did_end {
				continue
			}

			first := *container[(line_idx+1)%4]
			second := *container[(line_idx+2)%4]
			third := *container[(line_idx+3)%4]

			// Check Down
			if first[i] == check_runes[0] &&
				second[i] == check_runes[1] &&
				third[i] == check_runes[2] {
				accumilator++
			}

			// Check Down Righ
			if check_right &&
				first[i+1] == check_runes[0] &&
				second[i+2] == check_runes[1] &&
				third[i+3] == check_runes[2] {
				accumilator++
			}

			if i < 3 {
				continue
			}

			// Check Down Left
			if first[i-1] == check_runes[0] &&
				second[i-2] == check_runes[1] &&
				third[i-3] == check_runes[2] {
				accumilator++
			}
		}

		if !did_end {
			new_line_data := getNextLine(scanner)

			if new_line_data == nil {
				did_end = true
			} else {
				container[line_idx] = new_line_data
			}

		}

		line_idx = (line_idx + 1) % 4
	}

	return 4, strconv.Itoa(accumilator), "impl"
}

func getNextLine(scanner *bufio.Scanner) *[]rune {
	has_more := scanner.Scan()

	if !has_more {
		return nil
	}

	line := []rune(scanner.Text())

	return &line
}
