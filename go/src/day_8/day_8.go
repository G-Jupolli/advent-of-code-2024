package day8

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"strconv"
)

/*
Strategy
*/

type Point struct {
	x int
	y int
}

func DoDay8() (int, string, string) {

	file := helpers.GetFile(8)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part_1 := 0
	part_2 := -1

	var antinode_data []uint64

	var node_mapper = make(map[rune][]Point)

	line_count := 0
	line_len := 0

	for scanner.Scan() {
		line := scanner.Text()

		antinode_data = append(antinode_data, 0)

		for idx, char := range line {
			if char == '.' {
				continue
			}

			node_mapper[char] = append(node_mapper[char], Point{x: idx, y: line_count})
		}

		line_len = len(line) - 1
		line_count += 1
	}
	line_count -= 1

	for _, v := range node_mapper {
		for i := 0; i < len(v); i += 1 {
			left_node := v[i]

			for j := i + 1; j < len(v); j += 1 {
				right_node := v[j]

				a, b := findAntiNodes(left_node, right_node)

				if validateAntiNode(a, line_len, line_count) {
					x_flag := 1 << a.x

					if antinode_data[a.y]&uint64(x_flag) == 0 {
						antinode_data[a.y] |= uint64(x_flag)

						part_1 += 1
					}

				}

				if validateAntiNode(b, line_len, line_count) {
					x_flag := 1 << b.x

					if antinode_data[b.y]&uint64(x_flag) == 0 {
						antinode_data[b.y] |= uint64(x_flag)

						part_1 += 1
					}

				}
			}
		}
	}

	return 8, strconv.Itoa(part_1), strconv.Itoa(part_2)
}

func findAntiNodes(a Point, b Point) (Point, Point) {
	y_dif := a.y - b.y

	x_dif := a.x - b.x

	return Point{x: a.x + x_dif, y: a.y + y_dif}, Point{x: b.x - x_dif, y: b.y - y_dif}
}

func validateAntiNode(node Point, x_max int, y_max int) bool {
	if node.x < 0 ||
		node.y < 0 ||
		node.x > x_max ||
		node.y > y_max {
		return false
	}

	return true
}
