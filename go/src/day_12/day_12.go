package day12

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"strconv"
)

type GardenData struct {
	scanner        *bufio.Scanner
	reached_bottom bool
	map_height     int
	map_data       map[int][]rune

	part_1 int
}

const VISITED rune = '.'

func DoDay12() (int, string, string) {
	part_2 := 0

	file := helpers.GetFile(12)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	garden := GardenData{
		scanner,
		false,
		0,
		make(map[int][]rune),
		0,
	}

	i := 0

	for garden.checkRow(i) {

		// All of the plots in this row will be checked by now so can just delete here
		delete(garden.map_data, i)

		i += 1
	}

	return 12, strconv.Itoa(garden.part_1), strconv.Itoa(part_2)
}

func (g *GardenData) checkRow(row_idx int) bool {
	if row_idx == g.map_height && (g.reached_bottom || !g.pullNextRow()) {
		return false
	}

	for x, crop := range g.map_data[row_idx] {
		if crop == VISITED {
			continue
		}

		checked_cache := make(map[[2]int]bool)

		a, p := g.checkPlot(&checked_cache, crop, x, row_idx)

		g.part_1 += a * p
	}

	return true
}

func (g *GardenData) pullNextRow() bool {
	if !g.scanner.Scan() {
		g.reached_bottom = true
		return false
	}

	raw_bytes := []rune(g.scanner.Text())

	g.map_data[g.map_height] = raw_bytes

	g.map_height += 1

	return true
}

func (g *GardenData) checkPlot(checked_cache *map[[2]int]bool, checker rune, x int, y int) (int, int) {
	key := [2]int{x, y}

	if _, checked := (*checked_cache)[key]; checked {
		return 0, 0
	}

	if g.map_data[y][x] != checker {
		return 0, 1
	}

	(*checked_cache)[key] = true
	g.map_data[y][x] = VISITED

	a := 1
	p := 0

	// Check Left
	if x == 0 {
		p += 1
	} else {
		left_a, left_p := g.checkPlot(checked_cache, checker, x-1, y)
		a += left_a
		p += left_p
	}

	// Check Right
	if x == len(g.map_data[y])-1 {
		p += 1
	} else {
		right_a, right_p := g.checkPlot(checked_cache, checker, x+1, y)
		a += right_a
		p += right_p
	}

	// Check Up
	if _, has_up := g.map_data[y-1]; has_up {
		above_a, above_p := g.checkPlot(checked_cache, checker, x, y-1)
		a += above_a
		p += above_p
	} else {
		p += 1
	}

	// Check Down
	if y < g.map_height-1 {
		below_a, below_p := g.checkPlot(checked_cache, checker, x, y+1)
		a += below_a
		p += below_p
	} else if g.reached_bottom {
		p += 1
	} else if g.pullNextRow() {
		below_a, below_p := g.checkPlot(checked_cache, checker, x, y+1)
		a += below_a
		p += below_p
	} else {
		p += 1
	}

	return a, p
}
