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
	is_file := true
	file_id := 0
	memory := []byte{}

	for cursor < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		cursor++

		if is_file {
			for i := 0; i < int(char[0])-48; i++ {
				memory = append(memory, byte(file_id+48))
			}
			file_id++
			is_file = false
		} else {
			for i := 0; i < int(char[0])-48; i++ {
				memory = append(memory, '.')
			}
			is_file = true
		}
	}

	i := 0
	j := len(memory) - 1

	for i < j {
		if memory[i] == '.' && memory[j] >= 48 && memory[j] <= 57 {
			temp := memory[i]
			memory[i] = memory[j]
			memory[j] = temp
			i++
			j--
		} else if memory[i] == '.' {
			j--
		} else if memory[j] >= 48 && memory[j] <= 57 {
			i++
		} else {
			i++
			j--
		}
	}

	checksum := big.NewInt(0)
	for i := range memory {
		if memory[i] == '.' {
			break
		}
		file_idx := big.NewInt(int64(memory[i]) - 48)
		idx := big.NewInt(int64(i))
		num := new(big.Int).Mul(file_idx, idx)
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
