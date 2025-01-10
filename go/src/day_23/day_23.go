package day23

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"strconv"
)

/*

A computer in this case can be represented in a uint16 where the
2 bytes take up all the bits

This also means that a connection can be represented by a uint32 with
each 16 bit size computer taking up each side of the number

e.g. Computers aa bb cc

Connection List

aa-bb
bb-cc
aa-cc

When saving connections save bith aabb & bbaa

In the first connection, we register connection aa-bb in a hash map
We also denote aa && bb as known computers.

Second connection, register in map
Denote cc as known

Third connection
aa && cc are both known
For all known computers uu , check if uu-aa & uu-cc exists

If computer uu exists, we need to check if any of the computers start with a t
Where computer val >> 8 == 't'
*/

// We only care about the keys here so the values can be 0 size
type KnownComputers map[uint16]struct{}
type Connections map[uint32]struct{}

type Cache struct {
	cheif_groups    int
	known_computers KnownComputers
	connections     Connections
}

const DAY = 23

const search_char uint16 = 't'

func DoDay() (int, string, string) {
	file := helpers.GetFile(DAY)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// accumilator_first := 0
	accumilator_second := 0

	c := Cache{
		0,
		make(KnownComputers),
		make(Connections),
	}

	for scanner.Scan() {
		line := scanner.Bytes()

		a := parse_computer(line[0], line[1])
		b := parse_computer(line[3], line[4])

		c.register_connection(a, b)
	}

	return DAY, strconv.Itoa(c.cheif_groups), strconv.Itoa(accumilator_second)
}

func parse_computer(l byte, r byte) uint16 {
	return (uint16(l) << 8) | uint16(r)
}

func (c *Cache) search_3rd(a uint16, b uint16) {
	mask_a := uint32(a) << 16
	mask_b := uint32(b) << 16

	for u := range c.known_computers {
		if u == a || u == b {
			continue
		}

		mask_u := uint32(u)

		u_a := c.connections.has_connection(mask_a | mask_u)
		u_b := c.connections.has_connection(mask_b | mask_u)

		if u_a && u_b {
			c.check_cheif_group(a, b, u)
		}
	}
}

// Return true if both computers are already known
func (k KnownComputers) check_known(a uint16, b uint16) bool {
	_, has_a := k[a]
	_, has_b := k[b]

	if !has_a {
		k[a] = struct{}{}
	}
	if !has_b {
		k[b] = struct{}{}
	}

	return has_a && has_b
}

func (s Connections) has_connection(connection uint32) bool {
	_, ok := s[connection]

	return ok
}

func (c Cache) register_connection(a uint16, b uint16) {
	{ // Cache connection a -> b & b -> a
		mask_a := uint32(a)
		mask_b := uint32(b)

		atob := (mask_a << 16) | mask_b
		btoa := (mask_b << 16) | mask_a

		c.connections[atob] = struct{}{}
		c.connections[btoa] = struct{}{}
	}

	if c.known_computers.check_known(a, b) {
		c.search_3rd(a, b)
	}
}

func (cache *Cache) check_cheif_group(a uint16, b uint16, c uint16) {
	if (a>>8) == search_char || (b>>8) == search_char || (c>>8) == search_char {
		cache.cheif_groups += 1
	}
}
