package day24

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"fmt"
	"log"
	"strconv"
)

type bit_func byte

type register [3]byte

type cache struct {
	registers map[register]bool
	// Making the assumption z only ever gets to 64 bits
	// Looks like it doesn't go above 32 but I'll keep it as 64 just in cae
	z_val uint64
}

type process struct {
	reg_a register
	reg_b register
	reg_c register
	f     bit_func
}

const f_and bit_func = 0
const f_or bit_func = 1
const f_xor bit_func = 2

const DAY = 24

func DoDay() (int, string, string) {
	file := helpers.GetFile(DAY)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	accumilator_second := 0

	c := cache{
		make(map[register]bool),
		0,
	}

	c.inititialise_registers(scanner)

	var process_list []process

	for {
		has_process, p := parse_process(scanner)

		if !has_process {
			break
		}

		if !c.handle_process(p) {
			process_list = append(process_list, p)
		}
	}

	for len(process_list) > 0 {
		process_list = c.handle_process_list(process_list)
	}

	return DAY, strconv.FormatUint(c.z_val, 10), strconv.Itoa(accumilator_second)
}

func parse_process(scanner *bufio.Scanner) (bool, process) {

	if !scanner.Scan() {
		return false, process{}
	}

	line := scanner.Bytes()

	reg_a := register{
		line[0],
		line[1],
		line[2],
	}

	var p_fn bit_func
	ptr := 8

	switch line[4] {
	case 'A':
		p_fn = f_and
	case 'X':
		p_fn = f_xor
	case 'O':
		p_fn = f_or
		ptr -= 1
	default:
		log.Fatal("process parse fail on ", string(line))
		p_fn = f_and
	}

	reg_b := register{
		line[ptr],
		line[ptr+1],
		line[ptr+2],
	}

	reg_c := register{
		line[ptr+7],
		line[ptr+8],
		line[ptr+9],
	}

	return true, process{
		reg_a,
		reg_b,
		reg_c,
		p_fn,
	}
}

func (c *cache) inititialise_registers(scanner *bufio.Scanner) {

	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) < 3 {
			return
		}

		register_key := register{line[0], line[1], line[2]}

		c.registers[register_key] = line[5] == '1'
	}
}

func (c *cache) handle_process(p process) bool {
	val_a, has_a := c.registers[p.reg_a]
	if !has_a {
		return false
	}
	val_b, has_b := c.registers[p.reg_b]
	if !has_b {
		return false
	}

	var res bool
	switch p.f {
	case f_and:
		res = val_a && val_b
	case f_or:
		res = val_a || val_b
	case f_xor:
		res = val_a != val_b
	default:
		panic(fmt.Sprintf("unexpected day24.bit_func: %#v", p.f))
	}

	if p.reg_c[0] == 'z' {
		c.update_z(p.reg_c, res)
	} else {
		c.registers[p.reg_c] = res
	}

	return true
}

func (c *cache) update_z(r register, val bool) {
	if !val {
		return
	}

	idx := ((int(r[1]) - 48) * 10) + (int(r[2]) - 48)

	c.z_val |= uint64(1 << idx)
}

func (c *cache) handle_process_list(l []process) []process {

	var process_list []process

	// Reverse through list here to flip order
	// Not needed but can free up some operations
	for i := len(l) - 1; i >= 0; i -= 1 {
		p := l[i]

		if !c.handle_process(p) {
			process_list = append(process_list, p)
		}
	}

	return process_list
}
