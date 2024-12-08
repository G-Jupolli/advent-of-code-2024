package main

import (
	day7 "advent_of_code_2024/src/day_7"
	"fmt"
	"log"
	"os"
	"strings"
)

type DayRes struct {
	Day     int
	MaxLen  int
	PartOne string
	PartTwo string
}

func main() {
	os.Setenv("FULL_DATA", "yes")

	var data []DayRes

	// data = append(data, parseDayStats(day1.DoDay1()))
	// data = append(data, parseDayStats(day2.DoDay2()))
	// data = append(data, parseDayStats(day3.DoDay3()))
	// data = append(data, parseDayStats(day4.DoDay4()))
	// data = append(data, parseDayStats(day5.DoDay5()))
	// data = append(data, parseDayStats(day6.DoDay6()))
	data = append(data, parseDayStats(day7.DoDay7()))

	if len(data) < 1 {
		return
	}

	max_len := 0

	for _, item := range data {
		if item.MaxLen > max_len {
			max_len = item.MaxLen
		}
	}

	printContainers(false, true, max_len)
	for idx, item := range data {
		if idx != 0 {
			printContainers(true, true, max_len)
		}
		printValues(item, max_len)
	}
	printContainers(true, false, max_len)
}

func parseDayStats(day int, part_1 string, part_2 string) DayRes {
	max_len := len(part_1)

	if len(part_2) > max_len {
		max_len = len(part_2)
	}

	return DayRes{
		Day:     day,
		MaxLen:  max_len,
		PartOne: part_1,
		PartTwo: part_2,
	}
}

func printContainers(has_above bool, has_below bool, val_len int) {
	var left_pipe string
	var middle_pipe string
	var right_pipe string

	switch true {
	case has_above && has_below:
		left_pipe = "├"
		middle_pipe = "┼"
		right_pipe = "┤"
	case has_above && !has_below:
		left_pipe = "└"
		middle_pipe = "┴"
		right_pipe = "┘"
	case !has_above && has_below:
		left_pipe = "┌"
		middle_pipe = "┬"
		right_pipe = "┐"
	case !has_above && !has_below:
		fallthrough
	default:
		// Probably a bit much to fatal log here but that's fine
		log.Fatal("Invalid Logging state")
	}

	fmt.Printf("%s%s%s%s%s\n", left_pipe, strings.Repeat("─", 15), middle_pipe, strings.Repeat("─", val_len+2), right_pipe)
}

func printValues(data DayRes, val_len int) {

	fmt.Printf("│ Day %2d Part 1 │ %*s │\n", data.Day, val_len, data.PartOne)
	fmt.Printf("│        Part 2 │ %*s │\n", val_len, data.PartTwo)
}
