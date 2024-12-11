package day10_24

import (
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

type point struct {
	x int
	y int
}

func Part1() {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	cursor, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	trail := [][]int{}
	row := []int{}

	for cursor < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		cursor++

		if string(char[0]) == "\n" {
			trail = append(trail, row)
			row = []int{}
			continue
		}

		row = append(row, int(char[0])-48)
	}

	score := 0
	for i, row := range trail {
		for j := range row {
			if row[j] == 0 {
				counter := 0
				visited := make(map[point]bool)
				hike(trail, point{j, i}, &counter, &visited)
				score += counter
			}
		}
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(score))
	shared.ErrCheck(err)
	out.Sync()
}

func hike(trail [][]int, curr point, counter *int, visited *map[point]bool) {
	temp_visited := *visited
	if trail[curr.y][curr.x] == 9 {
		_, ok := temp_visited[curr]
		if !ok {
			*counter++
			temp_visited[curr] = true
			*visited = temp_visited
		}
		return
	}

	if curr.y-1 >= 0 && trail[curr.y-1][curr.x] == trail[curr.y][curr.x]+1 {
		hike(trail, point{curr.x, curr.y - 1}, counter, visited)
	}
	if curr.x+1 < len(trail[0]) && trail[curr.y][curr.x+1] == trail[curr.y][curr.x]+1 {
		hike(trail, point{curr.x + 1, curr.y}, counter, visited)
	}
	if curr.y+1 < len(trail) && trail[curr.y+1][curr.x] == trail[curr.y][curr.x]+1 {
		hike(trail, point{curr.x, curr.y + 1}, counter, visited)
	}
	if curr.x-1 >= 0 && trail[curr.y][curr.x-1] == trail[curr.y][curr.x]+1 {
		hike(trail, point{curr.x - 1, curr.y}, counter, visited)
	}
}
