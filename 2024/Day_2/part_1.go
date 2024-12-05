package day2_24

import (
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part1() {
	data := GetData()
	num_of_safe := 0
	var increasing bool
	var direction_set bool
	safe := false

	for _, arr := range data {
		for i := 0; i < len(arr)-1; i++ {
			if !direction_set {
				if arr[i] > arr[i+1] {
					increasing = false
				} else if arr[i] < arr[i+1] {
					increasing = true
				} else {
					safe = false
					break
				}
				direction_set = true
			}

			if increasing {
				if arr[i] >= arr[i+1] {
					safe = false
					break
				}
				if arr[i]-arr[i+1] < -3 || arr[i]-arr[i+1] > -1 {
					safe = false
					break
				}
			} else {
				if arr[i] <= arr[i+1] {
					safe = false
					break
				}
				if arr[i]-arr[i+1] > 3 || arr[i]-arr[i+1] < 1 {
					safe = false
					break
				}
			}
			safe = true
		}
		if safe {
			num_of_safe++
			safe = false
		}
		direction_set = false
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(num_of_safe))
	shared.ErrCheck(err)
	out.Sync()
}

func GetData() [][]int {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	digit := make([]byte, 1)
	arr := []int{}
	out := [][]int{}
	var num int

	for i < eof {
		_, err := f.Read(digit)
		shared.ErrCheck(err)
		i++

		if string(digit[0]) == "\n" {
			arr = append(arr, num)
			num = 0
			out = append(out, arr)
			arr = []int{}
		} else if digit[0] != ' ' {
			num = (num * 10) + (int(digit[0]) - 48)
		} else {
			arr = append(arr, num)
			num = 0
		}
	}

	return out
}
