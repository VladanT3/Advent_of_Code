package day5_24

import (
	"io"
	"os"
	"strconv"

	"github.com/VladanT3/Advent_of_Code"
)

func Part1() {
	rules := GetRules()
	data := GetData()
	correct := true
	sum := 0

	for _, arr := range data {
		for j := range arr {
			after, ok := rules[arr[j]]
			for k := j + 1; k < len(arr); k++ {
				_, ok = after[arr[k]]
				if !ok {
					correct = false
					break
				}
			}
			if !correct {
				break
			}
		}
		if correct {
			sum += arr[(len(arr)-1)/2]
		}
		correct = true
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(sum))
	shared.ErrCheck(err)
	out.Sync()
}

func GetRules() map[int]map[int]bool {
	f, err := os.Open("rules.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	sequence_map := make(map[int]map[int]bool)
	num := make([]byte, 2)
	var num1 int
	var num2 int

	for i < eof {
		_, err = f.Read(num)
		shared.ErrCheck(err)
		i += 2

		num1 = ((int(num[0]) - 48) * 10) + (int(num[1]) - 48)

		_, err = f.Seek(1, io.SeekCurrent)
		shared.ErrCheck(err)
		i++

		_, err = f.Read(num)
		shared.ErrCheck(err)
		i += 2

		num2 = ((int(num[0]) - 48) * 10) + (int(num[1]) - 48)

		_, err = f.Seek(1, io.SeekCurrent)
		shared.ErrCheck(err)
		i++

		if sequence_map[num1] == nil {
			sequence_map[num1] = make(map[int]bool)
		}
		sequence_map[num1][num2] = true
	}

	return sequence_map
}

func GetData() [][]int {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	num := 0
	out := [][]int{}
	arr := []int{}

	for i < eof {
		_, err := f.Read(char)
		shared.ErrCheck(err)
		i++

		if string(char[0]) == "\n" {
			arr = append(arr, num)
			num = 0
			out = append(out, arr)
			arr = []int{}
			continue
		}

		if char[0] == ',' {
			arr = append(arr, num)
			num = 0
		} else {
			num = num*10 + (int(char[0]) - 48)
		}
	}

	return out
}
