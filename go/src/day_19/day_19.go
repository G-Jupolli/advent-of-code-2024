package day19

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"strconv"
)

const DAY = 19

const WHITE = 'w'
const BLUE = 'u'
const BLACK = 'b'
const RED = 'r'
const GREEN = 'g'

type Towel []byte

type Pattern struct {
	pattern_ptr   int
	raw_pattern   []byte
	towel_indexes []int
}

func DoDay() (int, string, string) {
	file := helpers.GetFile(DAY)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	accumilator_first := 0
	accumilator_second := 0

	scanner.Scan()

	var towels []Towel

	var towel_buff []byte

	for _, c := range scanner.Bytes() {
		switch c {
		case ' ':
			continue
		case ',':
			towels = append(towels, towel_buff)
			towel_buff = nil
		default:
			towel_buff = append(towel_buff, c)
		}
	}

	towels = append(towels, towel_buff)

	scanner.Scan()

	for scanner.Scan() {
		var t_list []int
		p := Pattern{
			0,
			scanner.Bytes(),
			t_list,
		}

		if p.handle_pattern(towels) {
			accumilator_first += 1
		} else {
		}
	}

	return DAY, strconv.Itoa(accumilator_first), strconv.Itoa(accumilator_second)
}

func (p *Pattern) handle_pattern(towels []Towel) bool {

	start_idx := 0

	for {

		if p.pattern_ptr == len(p.raw_pattern) {
			return true
		}

		if p.insert_towel(towels, start_idx) {
			start_idx = 0
			continue
		}

		if len(p.towel_indexes) == 0 {
			return false
		}

		start_idx = p.towel_indexes[len(p.towel_indexes)-1]
		p.pattern_ptr -= len(towels[start_idx])
		p.towel_indexes = p.towel_indexes[:len(p.towel_indexes)-1]

		start_idx += 1
	}
}

func (p *Pattern) insert_towel(towels []Towel, start_idx int) bool {
	need_chr := p.raw_pattern[p.pattern_ptr]
	idx := start_idx

towel_loop:
	for i := idx; i < len(towels); i += 1 {
		t := towels[i]
		if t[0] == need_chr {
			// If adding the towel would overflow the space, skip
			if p.pattern_ptr+len(t) > len(p.raw_pattern) {
				continue towel_loop
			}

			for j := 1; j < len(t); j += 1 {
				if t[j] != p.raw_pattern[p.pattern_ptr+j] {
					continue towel_loop
				}
			}

			p.towel_indexes = append(p.towel_indexes, i)
			p.pattern_ptr += len(t)
			return true
		}
	}

	return false
}
