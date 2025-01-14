package day23

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"bytes"
	"sort"
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

	largest_group := c.find_largest_group()

	sort.Slice(largest_group, func(i, j int) bool {
		return largest_group[i] < largest_group[j]
	})

	return DAY, strconv.Itoa(c.cheif_groups), write_computers_to_string(largest_group)
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

func (c *Cache) register_connection(a uint16, b uint16) {
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

func (c *Cache) find_largest_group() []uint32 {

	var largest_group []uint32

	for {
		has_group, new_group := c.rip_group()

		if !has_group {
			break
		}

		if largest_group == nil || len(new_group) > len(largest_group) {
			largest_group = new_group
		}
	}

	return largest_group
}

func (c *Cache) rip_group() (bool, []uint32) {
	has, com := c.known_computers.rip_one()

	start_connection := uint32(com)

	if !has {
		return false, nil
	}

	group := []uint32{start_connection}

	var additional [][2]uint32

main_loop:
	for u := range c.known_computers {
		mask_u := uint32(u) << 16

		for i, g_c := range group {
			is_conn := c.connections.has_connection(mask_u | g_c)

			if !is_conn {
				if i != 0 {
					additional = append(additional, [2]uint32{uint32(u), g_c})
				}

				continue main_loop
			}
		}

		group = append(group, mask_u>>16)
	}

	if len(additional) == 0 {
		return true, group
	}

	for _, init := range additional {
		test_group := []uint32{start_connection, init[0]}

	addr_loop:
		for u := range c.known_computers {
			mask_u := uint32(u)

			if mask_u == init[0] || mask_u == init[1] {
				continue addr_loop
			}

			mask_u <<= 16

			for _, g_c := range test_group {
				if !c.connections.has_connection(mask_u | g_c) {
					continue addr_loop
				}
			}

			test_group = append(test_group, mask_u>>16)
		}

		if len(test_group) > len(group) {
			group = test_group
		}
	}

	return true, group
}

func (k KnownComputers) rip_one() (bool, uint16) {
	for c := range k {
		delete(k, c)

		return true, c
	}

	return false, 0
}

func write_computers_to_string(coms []uint32) string {

	var buffer bytes.Buffer

	for i, c := range coms {
		if i > 0 {
			buffer.WriteRune(',')
		}
		buffer.WriteByte(byte(c >> 8))
		buffer.WriteByte(byte(c))
	}

	return buffer.String()
}
func write_computers_to_string_16(coms []uint16) string {

	var buffer bytes.Buffer

	for i, c := range coms {
		if i > 0 {
			buffer.WriteRune(',')
		}
		buffer.WriteByte(byte(c >> 8))
		buffer.WriteByte(byte(c))
	}

	return buffer.String()
}
