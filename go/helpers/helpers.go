package helpers

import (
	"fmt"
	"log"
	"os"
)

type DayRes struct {
	Day     int
	PartOne string
	PartTwo string
}

func GetFile(day int, is_main bool) *os.File {

	var file_part string

	if is_main {
		file_part = "data"
	} else {
		file_part = "small"
	}

	path := fmt.Sprintf("../resources/day_%v_%v.txt", day, file_part)

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
