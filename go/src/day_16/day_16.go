package day16

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

/*

To do this we will do an implementation of Dijkstra's alforithm.

We need to make a queue of nodes and we only traverse the node with the current smallest weight.

When we split, we mark down where the next split node is and we try to continue to go forward.

If we cannot go forward or the current node we are looking at stops being the lowest weight.
We switch to traversing the lowest weight node until we reach the end.
*/

type maze_row [5]uint32

type maze_data struct {
	end      point
	paths    []maze_row
	visited  []maze_row
	travel_q []travel_node
	min_dist uint32
}

type point struct {
	x uint32
	y uint32
}

type travel_node struct {
	x        uint32
	y        uint32
	weight   int
	curr_dir heading
}

type heading int

const up heading = 0
const right heading = 1
const down heading = 2
const left heading = 3

const DAY = 16

const path = '.'
const wall = '#'
const start_char = 'S'
const end_char = 'E'

const turn_cost = 1_000

func DoDay() (int, string, string) {
	file := helpers.GetFile(DAY)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	accumilator_first := 0
	accumilator_second := 0

	m := parse_maze(scanner)

	if len(m.travel_q) == 0 {
		panic("should be at least 1 travel node")
	}

	curr_head := m.travel_q[0]
	m.travel_q = m.travel_q[1:]
	for {
		is_end, has_more := m.travel_node(&curr_head)

		if is_end {
			accumilator_first = curr_head.weight
			break
		}

		if !has_more {
			if len(m.travel_q) == 0 {
				panic("expected node to exist to switch to")
			}

			curr_head = m.travel_q[0]
			m.travel_q = m.travel_q[1:]
		}
	}

	return DAY, strconv.Itoa(accumilator_first), strconv.Itoa(accumilator_second)
}

func parse_maze(scanner *bufio.Scanner) maze_data {
	var paths []maze_row
	var visited []maze_row

	var start point
	var end point

	for scanner.Scan() {
		curr_path := maze_row{}

		line := scanner.Bytes()

		for i, b := range line {

			p_idx := i / 32

			switch b {
			case wall:
				continue
			case start_char:
				start = point{
					uint32(i),
					uint32(len(paths)),
				}
			case end_char:
				end = point{
					uint32(i),
					uint32(len(paths)),
				}
				fallthrough
			case path:
				curr_path[p_idx] |= uint32(1 << (i % 32))
			}

		}

		paths = append(paths, curr_path)
		visited = append(visited, maze_row{})
	}

	data := maze_data{
		end,
		paths,
		visited,
		[]travel_node{},
		^uint32(0),
	}

	data.initialise_travel_q(start.x, start.y)

	return data
}

func next_cords(x uint32, y uint32, dir heading) (uint32, uint32) {
	switch dir {
	case down:
		return x, y + 1
	case left:
		return x - 1, y
	case right:
		return x + 1, y
	case up:
		return x, y - 1
	default:
		panic(fmt.Sprintf("unexpected day16.heading: %#v", dir))
	}
}

func (m *maze_data) check_node(x uint32, y uint32) bool {

	block := x / 32
	flag := uint32(1 << (x % 32))

	if m.paths[y][block]&flag == 0 {
		return false
	}

	if m.visited[y][block]&flag != 0 {
		return false
	}

	return true
}

func (m *maze_data) mark_visited(x uint32, y uint32) bool {
	if x == m.end.x && y == m.end.y {
		return true
	}

	m.visited[y][x/32] |= uint32(1 << (x % 32))

	return false
}

func (m *maze_data) initialise_travel_q(x uint32, y uint32) {

	/*
		Start by looking to the right.
		need to include the cost when intialising routes.
	*/
	if m.check_node(x, y-1) {
		m.travel_q = append(m.travel_q, travel_node{
			x,
			y - 1,
			1 + turn_cost,
			up,
		})
	}
	if m.check_node(x+1, y) {
		m.travel_q = append(m.travel_q, travel_node{
			x + 1,
			y,
			1,
			right,
		})
	}
	if m.check_node(x, y+1) {
		m.travel_q = append(m.travel_q, travel_node{
			x,
			y + 1,
			1 + turn_cost,
			down,
		})
	}
	if m.check_node(x-1, y) {
		m.travel_q = append(m.travel_q, travel_node{
			x - 1,
			y,
			1 + turn_cost + turn_cost,
			left,
		})
	}
}

func (m *maze_data) register_turn(new_node travel_node) {
	if len(m.travel_q) == 0 {
		m.travel_q = []travel_node{new_node}
		return
	}

	i := sort.Search(
		len(m.travel_q),
		func(i int) bool {
			return m.travel_q[i].weight >= new_node.weight
		},
	)

	m.travel_q = append(m.travel_q, travel_node{})
	copy(m.travel_q[i+1:], m.travel_q[i:])
	m.travel_q[i] = new_node
}

func (m *maze_data) travel_node(node *travel_node) (bool, bool) {
	if !m.check_node(node.x, node.y) {
		return false, false
	}

	is_end := m.mark_visited(node.x, node.y)

	if is_end {
		return true, true
	}

	node.weight += 1

	{ // Check 90 deg turn
		next_dir := (node.curr_dir + 1) % 4
		next_x, next_y := next_cords(node.x, node.y, next_dir)

		if m.check_node(next_x, next_y) {
			m.register_turn(travel_node{
				next_x,
				next_y,
				node.weight + turn_cost,
				next_dir,
			})
		}
	}

	{ // Check 270 deg turn
		next_dir := (node.curr_dir + 3) % 4
		next_x, next_y := next_cords(node.x, node.y, next_dir)

		if m.check_node(next_x, next_y) {
			m.register_turn(travel_node{
				next_x,
				next_y,
				node.weight + turn_cost,
				next_dir,
			})
		}
	}

	{
		next_x, next_y := node.next_point()

		has_next := m.check_node(next_x, next_y)

		// This node can go no more
		// Can now be deleted
		if !has_next {
			return false, false
		}

		node.x = next_x
		node.y = next_y
	}

	if len(m.travel_q) == 0 {
		return false, true
	}

	{
		if m.travel_q[0].weight < node.weight {
			tmp := m.travel_q[0]

			if len(m.travel_q) == 1 {
				m.travel_q[0] = *node
				*node = tmp
				return false, true
			}

			i := sort.Search(
				len(m.travel_q),
				func(i int) bool {
					return m.travel_q[i].weight >= node.weight
				},
			)

			i -= 1

			copy(m.travel_q[:i], m.travel_q[1:i+1])
			m.travel_q[i] = *node
			*node = tmp
		}
	}

	return false, true
}

func (n *travel_node) next_point() (uint32, uint32) {
	return next_cords(n.x, n.y, n.curr_dir)
}
