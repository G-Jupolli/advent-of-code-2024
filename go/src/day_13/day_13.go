package day13

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

const part_2_scale = 10_000_000_000_000
const regex_pattern = `(\d+)(?:.*?)(\d+)`

type ArcadeState struct {
	scanner *bufio.Scanner
	re      *regexp.Regexp

	has_more bool
}

type Point struct {
	x float64
	y float64
}

type Game struct {
	a      Point
	b      Point
	target Point
}

func DoDay13() (int, string, string) {

	var re = regexp.MustCompile(regex_pattern)
	file := helpers.GetFile(13)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := ArcadeState{
		scanner,
		re,
		true,
	}

	part_1 := 0
	part_2 := 0

	for {
		has_game, game := state.getGame()

		if !has_game {
			break
		}

		has_solution, val := game.handleGame()

		if has_solution {
			part_1 += (val[0] * 3) + val[1]
		}

		new_target := Point{game.target.x + part_2_scale, game.target.y + part_2_scale}

		game.target = new_target

		has_solution, val = game.handleGame()

		if has_solution {
			part_2 += (val[0] * 3) + val[1]
		}
	}

	return 13, strconv.Itoa(part_1), strconv.Itoa(part_2)
}

func (s *ArcadeState) getGame() (bool, *Game) {
	if !s.has_more {
		return false, nil
	}

	a_data, err := s.parseLine()
	if err != nil {
		log.Fatalf("Failed to parse line %v", err)
	}
	b_data, err := s.parseLine()
	if err != nil {
		log.Fatalf("Failed to parse line %v", err)
	}
	target_data, err := s.parseLine()
	if err != nil {
		log.Fatalf("Failed to parse line %v", err)
	}

	game := Game{
		a_data,
		b_data,
		target_data,
	}

	s.has_more = s.scanner.Scan()

	return true, &game
}

func (s *ArcadeState) parseLine() (Point, error) {
	if !s.scanner.Scan() {
		log.Fatal("Should be line here")
	}

	matches := s.re.FindStringSubmatch(s.scanner.Text())

	if len(matches) < 3 {
		return Point{}, fmt.Errorf("invalid match count : %d", len(matches))
	}

	x, err := strconv.Atoi(matches[1])
	if err != nil {
		return Point{}, err
	}

	y, err := strconv.Atoi(matches[2])
	if err != nil {
		return Point{}, err
	}

	return Point{
		float64(x),
		float64(y),
	}, nil
}

func (g *Game) handleGame() (bool, [2]int) {

	determinant := g.a.x*g.b.y - g.a.y*g.b.x

	// When there are no solutions
	if determinant == 0 {
		return false, [2]int{}
	}

	x := (g.target.x*g.b.y - g.b.x*g.target.y) / determinant

	// When x is not an int
	// We cannot press a button 1.5 times
	if math.Mod(x, 1.0) != 0 {
		return false, [2]int{}
	}

	y := (g.a.x*g.target.y - g.target.x*g.a.y) / determinant

	return true, [2]int{int(x), int(y)}
}
