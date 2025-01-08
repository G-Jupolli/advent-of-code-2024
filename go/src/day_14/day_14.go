package day14

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Robot struct {
	start    Point
	velocity Point
}

const small_data_width = 11
const small_data_length = 7
const data_width = 101
const data_length = 103

const walk_steps = 100

const regex_pattern = `p=(.*?),(.*?) v=(.*?),(.*?)$`

func DoDay14() (int, string, string) {
	var bounds Point

	if os.Getenv("FULL_DATA") == "yes" {
		bounds = Point{data_width, data_length}
	} else {
		bounds = Point{small_data_width, small_data_length}
	}

	bounds_mid := Point{
		bounds.x / 2,
		bounds.y / 2,
	}

	var re = regexp.MustCompile(regex_pattern)

	file := helpers.GetFile(14)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	accumilator_second := 0

	robot_count := [4]int{0, 0, 0, 0}

	for scanner.Scan() {
		robot, err := parseRobot(re, scanner.Text())

		if err != nil {
			log.Fatal("Bad err", err)
		}

		// fmt.Printf("Robot : %v\n", robot)

		walk_point := robot.walk(bounds)

		// fmt.Printf("Walks to : %v\n", walk_point)

		if walk_point.x == bounds_mid.x || walk_point.y == bounds_mid.y {
			continue
		}

		quadrant := 0

		if walk_point.x > bounds_mid.x {
			quadrant += 1
		}

		if walk_point.y > bounds_mid.y {
			quadrant += 2
		}

		robot_count[quadrant] += 1
	}

	part_1 := robot_count[0] * robot_count[1] * robot_count[2] * robot_count[3]

	return 14, strconv.Itoa(part_1), strconv.Itoa(accumilator_second)
}

func parseRobot(re *regexp.Regexp, line string) (Robot, error) {
	matches := re.FindStringSubmatch(line)

	if len(matches) != 5 {
		return Robot{}, fmt.Errorf("invalid match count : %d", len(matches))
	}

	p_x, err := strconv.Atoi(matches[1])
	if err != nil {
		return Robot{}, err
	}
	p_y, err := strconv.Atoi(matches[2])
	if err != nil {
		return Robot{}, err
	}
	v_x, err := strconv.Atoi(matches[3])
	if err != nil {
		return Robot{}, err
	}
	v_y, err := strconv.Atoi(matches[4])
	if err != nil {
		return Robot{}, err
	}

	return Robot{
		start:    Point{p_x, p_y},
		velocity: Point{v_x, v_y},
	}, nil
}

/*

To find how far a robot walked we do:

x = (start.x + (velocity.x * 100)) % bounds.x

Same for y
*/

func (r *Robot) walk(bounds Point) Point {

	x := (r.start.x + (r.velocity.x * walk_steps)) % bounds.x

	if x < 0 {
		x += bounds.x
	}

	y := (r.start.y + (r.velocity.y * walk_steps)) % bounds.y

	if y < 0 {
		y += bounds.y
	}
	return Point{x, y}
}
