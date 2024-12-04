package day3_24

import (
	"io"
	"os"
	"strconv"
)

func Part1() {
	f, err := os.Open("input.txt")
	ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	ErrCheck(err)

	char := make([]byte, 1)
	mul_str := make([]byte, 3)
	found_mul := false
	num_1 := 0
	set_num_1 := false
	num_2 := 0
	mul_sum := 0

	for i < eof {
		_, err := f.Read(char)
		ErrCheck(err)
		i++

		if !found_mul {
			if char[0] == 'm' {
				_, err = f.Read(mul_str)
				ErrCheck(err)

				if string(mul_str) == "ul(" {
					found_mul = true
					i += 3
				} else {
					_, err = f.Seek(-3, io.SeekCurrent)
					ErrCheck(err)
				}
			}
		} else {
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

	out, err := os.Create("output.txt")
	ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(mul_sum))
	ErrCheck(err)
	out.Sync()
}

func ErrCheck(e error) {
	if e != nil {
		panic(e)
	}
}
