package day10

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"strconv"
)

type TravelData struct {
	head_count int
	rating     int
	visited    map[[2]int]bool
}

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

		data := TravelData{
			0,
			0,
			make(map[[2]int]bool),
		}

		data.travel(map_data, 0, start_point[0], start_point[1])

		part_1 += data.head_count
		part_2 += data.rating
	}

	return 10, strconv.Itoa(part_1), strconv.Itoa(part_2)
}

func (d *TravelData) travel(map_data [][]int, curr_val int, x int, y int) {
	if curr_val == 9 {
		key := [2]int{x, y}

		d.rating += 1

		if _, visited := d.visited[key]; visited {
			return
		}

		d.head_count += 1
		return
	}

	if y > 0 && map_data[y-1][x] == curr_val+1 {
		d.travel(map_data, curr_val+1, x, y-1)
	}

	if y < len(map_data)-1 && map_data[y+1][x] == curr_val+1 {
		d.travel(map_data, curr_val+1, x, y+1)
	}

	if x < len(map_data[y])-1 && map_data[y][x+1] == curr_val+1 {
		d.travel(map_data, curr_val+1, x+1, y)
	}

	if x > 0 && map_data[y][x-1] == curr_val+1 {
		d.travel(map_data, curr_val+1, x-1, y)
	}
}
