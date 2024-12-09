package day9_24

import (
	"fmt"
	"io"
	"math/big"
	"os"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part1() {
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

	i := 0
	j := len(memory) - 1

	for i < j {
		if memory[i] == -1 && memory[j] >= 0 {
			temp := memory[i]
			memory[i] = memory[j]
			memory[j] = temp
			i++
			j--
		} else if memory[i] == -1 {
			j--
		} else if memory[j] >= 0 {
			i++
		} else {
			i++
			j--
		}
	}

	for i := range memory {
		if memory[i] == -1 {
			break
		}
		file_id := big.NewInt(int64(memory[i]))
		idx := big.NewInt(int64(i))
		num := new(big.Int).Mul(file_id, idx)
		checksum = checksum.Add(checksum, num)
	}

	fmt.Println(checksum)

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + checksum.String())
	shared.ErrCheck(err)
	out.Sync()
}
