package day17

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"log"
	"strconv"
	"strings"
)

/*
The adv instruction (opcode 0) performs division.
The numerator is the value in the A register.
The denominator is found by raising 2 to the power of the instruction's combo operand.
(So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.)
The result of the division operation is truncated to an integer and then written to the A register.
*/
const OP_adv = 0

/*
The bxl instruction (opcode 1) calculates the bitwise XOR of register B
and the instruction's literal operand, then stores the result in register B.
*/
const OP_bxl = 1
const OP_bst = 2
const OP_jnz = 3
const OP_bxc = 4
const OP_out = 5
const OP_bdv = 6
const OP_cdv = 7

type Cache struct {
	instruction_pointer int
	register_a          int
	register_b          int
	register_c          int
	out                 []int
}

func DoDay17() (int, string, string) {
	file := helpers.GetFile(17)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	second_safe_reactors := 0

	cache, instruction_list := parse_input(scanner)

	for {
		if cache.instruction_pointer >= len(instruction_list) {
			break
		}

		op_code := instruction_list[cache.instruction_pointer]
		operand := instruction_list[cache.instruction_pointer+1]

		cache.handle_op_code(op_code, operand)
	}

	var sb strings.Builder
	for i, v := range cache.out {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(v))
	}

	return 17, sb.String(), strconv.Itoa(second_safe_reactors)
}

func parse_input(scanner *bufio.Scanner) (Cache, []int) {
	register_a := parse_register(scanner)
	register_b := parse_register(scanner)
	register_c := parse_register(scanner)

	scanner.Scan()
	scanner.Scan()

	program := scanner.Bytes()

	var instruction_list []int

	// First 9 bytes are 'Program: ' so we skip these
	// Every other byte is a comma so we skip these too
	for i := 9; i < len(program); i += 2 {
		// -48 to bring to int value of the rune
		instruction_list = append(instruction_list, int(program[i])-48)
	}

	var out_collector []int

	return Cache{
		0,
		register_a,
		register_b,
		register_c,
		out_collector,
	}, instruction_list
}

func parse_register(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		log.Fatal("Invalid Input")
		return -1
	}

	line := scanner.Text()

	val := line[12:]

	int_val, err := strconv.Atoi(val)

	if err != nil {
		log.Fatal("Should be int ", err)
		return -1
	}

	return int_val
}

func (c *Cache) get_operand_value(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.register_a
	case 5:
		return c.register_b
	case 6:
		return c.register_c
	default:
		log.Fatal("invalid operand ", operand)
		return -1
	}
}

func (c *Cache) handle_op_code(op_code int, operand int) {

	switch op_code {
	case OP_adv:
		// Numerator is register a
		numerator := c.register_a
		// Left shift 1 by operand value to find 2 to the power of operand value
		denominator := int(1 << c.get_operand_value(operand))
		// Load division into register a
		c.register_a = numerator / denominator
	case OP_bxl:
		// Load XOR of register b and operand literal into register b
		c.register_b ^= operand
	case OP_bst:
		// Load mod 8 of operand value into register b
		c.register_b = c.get_operand_value(operand) & 0b0111
	case OP_jnz:
		if c.register_a != 0 {
			c.instruction_pointer = operand
			return
		}
	case OP_bxc:
		c.register_b ^= c.register_c
	case OP_out:
		c.out = append(c.out, (c.get_operand_value(operand) & 0b0111))
	case OP_bdv:
		numerator := c.register_a
		denominator := int(1 << c.get_operand_value(operand))
		c.register_b = numerator / denominator
	case OP_cdv:
		numerator := c.register_a
		denominator := int(1 << c.get_operand_value(operand))
		c.register_c = numerator / denominator
	default:
		log.Fatal("invalid command", op_code)
	}

	c.instruction_pointer += 2
}
