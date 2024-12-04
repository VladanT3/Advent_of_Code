package day1_24

import (
	"io"
	"math"
	"os"
	"sort"
	"strconv"
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
	ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(sum))
	ErrCheck(err)

	out.Sync()
}

func ErrCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func GetLists() ([]int, []int) {
	f, err := os.Open("input.txt")
	ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	ErrCheck(err)

	arr1 := []int{}
	arr2 := []int{}
	str_num := make([]byte, 5)

	for i < eof {
		_, err := f.Read(str_num)
		ErrCheck(err)
		num, err := strconv.Atoi(string(str_num))
		ErrCheck(err)
		arr1 = append(arr1, num)

		_, err = f.Seek(3, io.SeekCurrent)
		ErrCheck(err)

		_, err = f.Read(str_num)
		ErrCheck(err)
		num, err = strconv.Atoi(string(str_num))
		ErrCheck(err)
		arr2 = append(arr2, num)

		_, err = f.Seek(1, io.SeekCurrent)
		ErrCheck(err)

		i += 14
	}

	return arr1, arr2
}
