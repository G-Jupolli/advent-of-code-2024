package day10

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"strconv"
)

func DoDay10() (int, string, string) {
	part_1 := 0
	part_2 := 0

	file := helpers.GetFile(10)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var map_data [][]int
	var start_points [][2]int

	for scanner.Scan() {
		var curr_row []int
		for i, char := range scanner.Bytes() {
			height := int(char - 48)

			if height == 0 {
				start_points = append(start_points, [2]int{i, len(map_data)})
			}

			curr_row = append(curr_row, height)
		}
		map_data = append(map_data, curr_row)
	}

	for _, start_point := range start_points {
		var visited [][2]int
		travel(map_data, 0, start_point[0], start_point[1], &part_1, &visited, &part_2)
	}

	return 10, strconv.Itoa(part_1), strconv.Itoa(part_2)
}

func travel(map_data [][]int, curr_val int, x int, y int, head_count *int, visited *[][2]int, rating *int) {
	if curr_val == 9 {
		*rating = *rating + 1
		for _, point := range *visited {
			if point[0] == x && point[1] == y {
				return
			}
		}
		*head_count = *head_count + 1
		*visited = append(*visited, [2]int{x, y})
		return
	}

	if y > 0 && map_data[y-1][x] == curr_val+1 {
		travel(map_data, curr_val+1, x, y-1, head_count, visited, rating)
	}

	if y < len(map_data)-1 && map_data[y+1][x] == curr_val+1 {
		travel(map_data, curr_val+1, x, y+1, head_count, visited, rating)
	}

	if x < len(map_data[y])-1 && map_data[y][x+1] == curr_val+1 {
		travel(map_data, curr_val+1, x+1, y, head_count, visited, rating)
	}

	if x > 0 && map_data[y][x-1] == curr_val+1 {
		travel(map_data, curr_val+1, x-1, y, head_count, visited, rating)
	}
}
