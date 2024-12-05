package day3_24

import (
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	mul_str := make([]byte, 3)
	dont_str := make([]byte, 6)
	do_str := make([]byte, 3)
	found_mul := false
	found_dont := false
	num_1 := 0
	set_num_1 := false
	num_2 := 0
	mul_sum := 0

	for i < eof {
		_, err := f.Read(char)
		shared.ErrCheck(err)
		i++

		if !found_dont {
			if char[0] == 'd' {
				_, err = f.Read(dont_str)
				shared.ErrCheck(err)

				if string(dont_str) == "on't()" {
					found_dont = true
					i += 6
				} else {
					_, err = f.Seek(-6, io.SeekCurrent)
					shared.ErrCheck(err)
				}
			}
		} else {
			if char[0] == 'd' {
				_, err = f.Read(do_str)
				shared.ErrCheck(err)

				if string(do_str) == "o()" {
					found_dont = false
					i += 3
				} else {
					_, err = f.Seek(-3, io.SeekCurrent)
					shared.ErrCheck(err)
				}
			}
		}

		if !found_mul && !found_dont {
			if char[0] == 'm' {
				_, err = f.Read(mul_str)
				shared.ErrCheck(err)

				if string(mul_str) == "ul(" {
					found_mul = true
					i += 3
				} else {
					_, err = f.Seek(-3, io.SeekCurrent)
					shared.ErrCheck(err)
				}
			}
		} else if found_mul && !found_dont {
			if !set_num_1 {
				if int(char[0]) >= 48 && int(char[0]) <= 57 {
					num_1 = (num_1 * 10) + (int(char[0]) - 48)
				} else if num_1 != 0 && char[0] == ',' {
					set_num_1 = true
				} else {
					found_mul = false
					num_1 = 0
				}
			} else {
				if int(char[0]) >= 48 && int(char[0]) <= 57 {
					num_2 = (num_2 * 10) + (int(char[0]) - 48)
				} else if num_2 != 0 && char[0] == ')' {
					mul_sum += num_1 * num_2
					found_mul = false
					num_1 = 0
					num_2 = 0
					set_num_1 = false
				} else {
					found_mul = false
					num_1 = 0
					num_2 = 0
					set_num_1 = false
				}
			}
		}
	}

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + strconv.Itoa(mul_sum))
	shared.ErrCheck(err)
	out.Sync()
}
