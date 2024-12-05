package day1_24

import (
	"io"
	"math"
	"os"
	"sort"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part1() {
	arr1, arr2 := GetLists()

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	sum := 0
	for i := range arr1 {
		diff := math.Abs(float64(arr1[i] - arr2[i]))
		sum += int(diff)
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(sum))
	shared.ErrCheck(err)

	out.Sync()
}

func GetLists() ([]int, []int) {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	arr1 := []int{}
	arr2 := []int{}
	str_num := make([]byte, 5)

	for i < eof {
		_, err := f.Read(str_num)
		shared.ErrCheck(err)
		num, err := strconv.Atoi(string(str_num))
		shared.ErrCheck(err)
		arr1 = append(arr1, num)

		_, err = f.Seek(3, io.SeekCurrent)
		shared.ErrCheck(err)

		_, err = f.Read(str_num)
		shared.ErrCheck(err)
		num, err = strconv.Atoi(string(str_num))
		shared.ErrCheck(err)
		arr2 = append(arr2, num)

		_, err = f.Seek(1, io.SeekCurrent)
		shared.ErrCheck(err)

		i += 14
	}

	return arr1, arr2
}
