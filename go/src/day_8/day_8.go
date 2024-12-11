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
	part_2 := 0

	var antinode_data []uint64
	var antinode_data_p2 []uint64

	var node_mapper = make(map[rune][]Point)

	line_count := 0
	line_len := 0

	for scanner.Scan() {
		line := scanner.Text()

		antinode_data = append(antinode_data, 0)
		antinode_data_p2 = append(antinode_data_p2, 0)

		for idx, char := range line {
			if char == '.' {
				continue
			}

			node_mapper[char] = append(node_mapper[char], Point{x: idx, y: line_count})
		}

		line_len = len(line)
		line_count += 1
	}

	for _, v := range node_mapper {
		for i, left_node := range v {
			for j := i + 1; j < len(v); j += 1 {
				right_node := v[j]

				for idx, anti_node := range findValidAntiNodes(left_node, right_node, line_len, line_count) {
					intrT := 1 << anti_node.x
					x_flag := uint64(intrT)

					// Only check for p1 on 1st idx
					if idx == 1 {
						if antinode_data[anti_node.y]&x_flag == 0 {
							antinode_data[anti_node.y] |= x_flag

							part_1 += 1
						}
					}

					if antinode_data_p2[anti_node.y]&x_flag == 0 {
						antinode_data_p2[anti_node.y] |= x_flag

						part_2 += 1
					}
				}

				for idx, anti_node := range findValidAntiNodes(right_node, left_node, line_len, line_count) {
					x_flag := uint64(1 << anti_node.x)

					// Only check for p1 on 0th idx
					if idx == 1 {
						if antinode_data[anti_node.y]&x_flag == 0 {
							antinode_data[anti_node.y] |= x_flag

							part_1 += 1
						}
					}

					if antinode_data_p2[anti_node.y]&x_flag == 0 {
						antinode_data_p2[anti_node.y] |= x_flag

						part_2 += 1
					}
				}
			}
		}
	}

	return 8, strconv.Itoa(part_1), strconv.Itoa(part_2)
}

func findValidAntiNodes(a Point, b Point, x_max int, y_max int) []Point {

	y_dif := a.y - b.y
	x_dif := a.x - b.x

	var node_list []Point

	node_list = append(node_list, a)

	curr_point := Point{
		x: a.x + x_dif,
		y: a.y + y_dif,
	}

	for validateAntiNode(curr_point, x_max, y_max) {
		node_list = append(node_list, curr_point)

		curr_point.x += x_dif
		curr_point.y += y_dif
	}

	return node_list
}

func validateAntiNode(node Point, x_max int, y_max int) bool {
	if node.x < 0 ||
		node.y < 0 ||
		node.x >= x_max ||
		node.y >= y_max {
		return false
	}

	return true
}
