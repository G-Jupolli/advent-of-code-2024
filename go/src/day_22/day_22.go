package day22

import (
	"advent_of_code_2024/helpers"
	"bufio"
	"log"
	"strconv"
)

/*
Hash process:
  - Multiply by 64, mix into secret then prune
  - Divide by 32, round to int, mix in then prune
  - Multiply by 2048, mix prune

Key:
  - Mix ~ XOR given and secret, set to secret
  - Prune ~ secret = secret % 16_777_216, 2^24
    can be secret & (2^24 - 1)
    (1 << 24) -1

Bit porcess:
Secret Value = s

a = s << 6
b = s ^ a
c = b & ( ( 1 << 24 ) - 1)

# The prune operator dictates that the numbers must be contained in 23 bits

If we have 2 pointers, 1 on the 23rd bit & one on the 17th bit (23 - 6)
This meets the 3rd condition as the number is limited to 23 bits
Pointers pa & pb
We then right shift the pointers until we bit pb < 0

The first 2 conditions are met by pa ^ pb

3rd process is the same with the pointers starting at 23 & 12

# Division by 32 is the same as right shifting by 5 bits

a = s >> 5
b = s ^ a

Pointer Starts:
  - pa ~ 23
  - pb ~ 17
  - pc ~ 28
  - pd ~ 12

Result is the bit at pa: ( ( pa ^ pb ) ^ pc ) ^ pd

	Equivalent to  : ( pa ^ pb ) ^ ( pc ^ pd )
*/
const pb_delta = 6
const pc_delta = 5
const pd_delta = 11

const prune_val = uint32(1<<24) - 1
const DAY = 22

func DoDay() (int, string, string) {
	file := helpers.GetFile(DAY)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	accumilator_first := 0
	accumilator_second := 0

	for scanner.Scan() {
		line := scanner.Text()

		val, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal("Invalid input ", line)
		}

		val_p := uint32(val)
		for i := 0; i < 2000; i += 1 {
			val_p = process_val(val_p)
		}

		// fmt.Println(val, ": ", val_p)

		accumilator_first += int(val_p)
	}

	return DAY, strconv.Itoa(accumilator_first), strconv.Itoa(accumilator_second)
}

func process_val(inp uint32) uint32 {

	a := (inp ^ (inp << 6)) & prune_val
	b := (a ^ (a >> 5)) & prune_val
	res := (b ^ (b << 11)) & prune_val

	return res
}

func _process_val(inp uint32) uint32 {
	result := uint32(0)

	// Loop is based on pa so can be initialised here
	for pa := 23; pa >= 0; pa -= 1 {

		bit_a := (inp >> pa) & 1
		bit_c := (inp >> (pa + pc_delta)) & 1

		var bit_b uint32
		if pa >= pb_delta {
			bit_b = (inp >> (pa - pb_delta)) & 1
		}

		var bit_d uint32
		if pa >= pd_delta {
			bit_b = (inp >> (pa - pd_delta)) & 1
		}
		res := (bit_a ^ bit_b) ^ (bit_c ^ bit_d)

		result |= (res << pa)
	}

	return result
}
