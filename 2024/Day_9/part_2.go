package day9_24

import (
	"fmt"
	"io"
	"math/big"
	"os"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	cursor, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	file_id := 0
	is_file := true
	memory := []int{}
	checksum := new(big.Int)

	for cursor < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		cursor++

		if is_file {
			for range int(char[0]) - 48 {
				memory = append(memory, file_id)
			}
			file_id++
			is_file = false
		} else {
			for range int(char[0]) - 48 {
				memory = append(memory, -1)
			}
			is_file = true
		}
	}

	file := -1
	length_of_file := 0
	file_start := -1
	for i := len(memory) - 1; i >= 0; i-- {
		if memory[i] >= 0 {
			file = memory[i]
			file_start = i
			break
		}
	}

	length_of_free_space := 0
	free_space_start := -1

	moved := make(map[int]bool)

	for i := len(memory) - 1; i >= 0; i-- {
		_, ok := moved[i]
		if ok {
			continue
		}
		if memory[i] != -1 {
			if file != memory[i] {
				for j := 0; j <= file_start-length_of_file; j++ {
					if memory[j] == -1 {
						if free_space_start == -1 {
							free_space_start = j
						}
						length_of_free_space++
						if length_of_free_space == length_of_file {
							for length_of_file > 0 {
								temp := memory[file_start]
								memory[file_start] = memory[free_space_start]
								memory[free_space_start] = temp
								moved[free_space_start] = true
								length_of_file--
								file_start--
								free_space_start++
							}
							length_of_free_space = 0
							free_space_start = -1
							break
						}
					} else {
						length_of_free_space = 0
						free_space_start = -1
					}
				}

				file_start = i
				file = memory[i]
				length_of_file = 1
			} else {
				length_of_file++
			}
		}
	}

	for i := range memory {
		if memory[i] == -1 {
			continue
		}
		file_id := big.NewInt(int64(memory[i]))
		idx := big.NewInt(int64(i))
		num := new(big.Int).Mul(file_id, idx)
		checksum = checksum.Add(checksum, num)
	}

	fmt.Println(checksum)

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + checksum.String())
	shared.ErrCheck(err)
	out.Sync()
}
