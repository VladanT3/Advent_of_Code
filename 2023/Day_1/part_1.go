package day1_23

import (
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part1() {
	data := GetData()
	sum := 0
	var num1 int
	var num2 int

	for _, arr := range data {
		for j := range arr {
			if int(arr[j]) >= 48 && int(arr[j]) <= 57 {
				num1 = int(arr[j]) - 48
				for k := len(arr) - 1; k >= j; k-- {
					if int(arr[k]) >= 48 && int(arr[k]) <= 57 {
						num2 = int(arr[k]) - 48
						break
					}
				}
				break
			}
		}
		sum += (num1 * 10) + num2
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(sum))
	out.Sync()
}

func GetData() [][]byte {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	out := [][]byte{}
	arr := []byte{}

	for i < eof {
		_, err := f.Read(char)
		shared.ErrCheck(err)
		i++

		if string(char[0]) == "\n" {
			out = append(out, arr)
			arr = []byte{}
			continue
		}

		arr = append(arr, char[0])
	}

	return out
}
