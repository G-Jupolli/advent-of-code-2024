package day12

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"log"
	"strconv"
)

type CachPointMap map[[2]int]bool

type GardenData struct {
	scanner        *bufio.Scanner
	reached_bottom bool
	map_height     int
	map_data       map[int][]rune

	part_1 int
	part_2 int

	cache_visited      CachPointMap
	cache_corners      CachPointMap
	cache_corner_count int
}

const VISITED rune = '.'

func DoDay12() (int, string, string) {

	file := helpers.GetFile(12)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	garden := GardenData{
		scanner,
		false,
		0,
		make(map[int][]rune),
		0,
		0,
		make(CachPointMap),
		make(CachPointMap),
		0,
	}

	i := 0

	for garden.checkRow(i) {

		// All of the plots in this row will be checked by now so can just delete here
		delete(garden.map_data, i)

		i += 1
	}

	return 12, strconv.Itoa(garden.part_1), strconv.Itoa(garden.part_2)
}

func (g *GardenData) checkRow(row_idx int) bool {
	if row_idx == g.map_height && (g.reached_bottom || !g.pullNextRow()) {
		return false
	}

	for x, crop := range g.map_data[row_idx] {
		if crop == VISITED {
			continue
		}

		_, a, p := g.checkPlot(crop, x, row_idx)

		// println(g.cache_corners)

		for corner := range g.cache_corners {
			g.checkInnerCorner(corner[0], corner[1])
		}

		if a > 0 && g.cache_corner_count < 4 {
			log.Fatal("FFS", crop)
		}
		println(a, "*", g.cache_corner_count, "=", a*g.cache_corner_count)
		g.part_1 += a * p
		g.part_2 += a * g.cache_corner_count

		g.resetCache()
	}

	return true
}

func (g *GardenData) resetCache() {
	g.cache_visited = make(CachPointMap)
	g.cache_corners = make(CachPointMap)
	g.cache_corner_count = 0
}

func (g *GardenData) pullNextRow() bool {
	if g.reached_bottom {
		return false
	}

	if !g.scanner.Scan() {
		g.reached_bottom = true
		return false
	}

	raw_bytes := []rune(g.scanner.Text())

	g.map_data[g.map_height] = raw_bytes

	g.map_height += 1

	return true
}

func (g *GardenData) checkPlot(checker rune, x int, y int) (bool, int, int) {
	key := [2]int{x, y}

	if _, checked := g.cache_visited[key]; checked {
		return true, 0, 0
	}

	if g.map_data[y][x] != checker {
		g.cache_corners[key] = true
		return false, 0, 1
	}

	g.cache_visited[key] = true
	g.map_data[y][x] = VISITED

	a := 1
	p := 0

	x_dir := 0
	sides := 0

	// 12 * 28 + 4* 4 + 4 * 4
	// Check Left
	if x > 0 {
		is_node, left_a, left_p := g.checkPlot(checker, x-1, y)
		a += left_a
		p += left_p

		if !is_node {
			x_dir += 1
		}
	} else {
		p += 1
		x_dir += 1

		g.cache_corners[[2]int{x - 1, y}] = true
	}

	// Check Right
	if x < len(g.map_data[y])-1 {
		is_node, right_a, right_p := g.checkPlot(checker, x+1, y)
		a += right_a
		p += right_p

		if !is_node {
			x_dir += 1
		}
	} else {
		p += 1
		x_dir += 1
		g.cache_corners[[2]int{x + 1, y}] = true
	}

	// Check Up
	if _, has_up := g.map_data[y-1]; has_up {
		is_node, above_a, above_p := g.checkPlot(checker, x, y-1)
		a += above_a
		p += above_p

		if !is_node {
			sides += x_dir
		}
	} else {
		p += 1
		sides += x_dir
		g.cache_corners[[2]int{x, y - 1}] = true
	}

	// Check Down
	if y < g.map_height-1 || g.pullNextRow() {
		is_node, below_a, below_p := g.checkPlot(checker, x, y+1)
		a += below_a
		p += below_p

		if !is_node {
			sides += x_dir
		}
	} else {
		p += 1
		sides += x_dir
		g.cache_corners[[2]int{x, y + 1}] = true
	}

	if sides > 0 {
		g.cache_corner_count += sides
	}

	return true, a, p
}

func (g *GardenData) checkInnerCorner(x int, y int) {
	// fmt.Printf("%v\n", g.cache_visited)

	x_dir := 0

	// Check Left
	// fmt.Println(x-1, y)

	if x > 0 && g.cache_visited[[2]int{x - 1, y}] {
		x_dir += 1
	}

	// fmt.Println(x+1, y)
	// Check Right
	if x < len(g.map_data[y])-1 && g.cache_visited[[2]int{x + 1, y}] {
		x_dir += 1
	}

	if x_dir == 0 {
		return
	}

	// fmt.Println(x, y-1)
	sides := 0
	// Check Up
	if _, has_up := g.map_data[y-1]; has_up && g.cache_visited[[2]int{x, y - 1}] {
		sides += x_dir
	}

	if y == g.map_height-1 {
		g.pullNextRow()
	}

	// fmt.Println(x, y+1)
	// Check Down
	if y < g.map_height-1 && g.cache_visited[[2]int{x, y + 1}] {
		sides += x_dir
	}

	g.cache_corner_count += sides

}
