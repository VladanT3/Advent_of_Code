package day6_24

import (
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part1() {
	level_map, start := GetMapAndStart()
	path_len := 0
	seen_path := [][]bool{}

	for _, arr := range level_map {
		temp_arr := []bool{}
		for range arr {
			temp_arr = append(temp_arr, false)
		}
		seen_path = append(seen_path, temp_arr)
	}

	i := start[0]
	j := start[1]
	path_found := false

	for {
		for level_map[i][j] != '#' {
			if !seen_path[i][j] {
				path_len++
				seen_path[i][j] = true
			}
			i--
			if i < 0 {
				path_found = true
				break
			}
		}
		if path_found {
			break
		}
		i++

		for level_map[i][j] != '#' {
			if !seen_path[i][j] {
				path_len++
				seen_path[i][j] = true
			}
			j++
			if j >= len(level_map[0]) {
				path_found = true
				break
			}
		}
		if path_found {
			break
		}
		j--

		for level_map[i][j] != '#' {
			if !seen_path[i][j] {
				path_len++
				seen_path[i][j] = true
			}
			i++
			if i >= len(level_map) {
				path_found = true
				break
			}
		}
		if path_found {
			break
		}
		i--

		for level_map[i][j] != '#' {
			if !seen_path[i][j] {
				path_len++
				seen_path[i][j] = true
			}
			j--
			if j < 0 {
				path_found = true
				break
			}
		}
		if path_found {
			break
		}
		j++
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(path_len))
	shared.ErrCheck(err)
	out.Sync()
}

func GetMapAndStart() ([][]byte, []int) {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	curr, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	out := [][]byte{}
	arr := []byte{}
	char := make([]byte, 1)
	i := 0
	j := 0
	stop_counting := false

	for curr < eof {
		_, err := f.Read(char)
		shared.ErrCheck(err)
		curr++

		if char[0] == '^' {
			stop_counting = true
		}

		if string(char[0]) == "\n" {
			if !stop_counting {
				i++
				j = 0
			}
			out = append(out, arr)
			arr = []byte{}
			continue
		}

		arr = append(arr, char[0])
		if !stop_counting {
			j++
		}
	}

	return out, []int{i, j}
}
