package day13_24

import (
	"fmt"
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part1() {
	data := GetData()
	a := 0
	b := 0
	tokens := 0

	for _, machine := range data {
		b = ((machine[4] * machine[1]) - (machine[5] * machine[0])) / ((machine[2] * machine[1]) - (machine[3] * machine[0]))
		a = (machine[4] - b*machine[2]) / machine[0]

		if machine[0]*a+machine[2]*b == machine[4] && machine[1]*a+machine[3]*b == machine[5] {
			tokens += a*3 + b
		}
	}

	fmt.Println(tokens)

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(tokens))
	shared.ErrCheck(err)
	out.Sync()
}

func GetData() [][]int {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	cursor, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	out := [][]int{}
	arr := []int{}
	num := 0

	for cursor < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		cursor++

		if int(char[0]) >= 48 && int(char[0]) <= 57 {
			num = (num * 10) + (int(char[0]) - 48)
		} else if num != 0 {
			arr = append(arr, num)
			num = 0
		}

		if len(arr) == 6 {
			out = append(out, arr)
			arr = []int{}
		}
	}

	return out
}
