package day_10_24

import (
	"fmt"
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
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
				hikePart2(trail, point{j, i}, &counter)
				score += counter
			}
		}
	}

	fmt.Println(score)

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(score))
	shared.ErrCheck(err)
	out.Sync()
}

func hikePart2(trail [][]int, curr point, counter *int) {
	if trail[curr.y][curr.x] == 9 {
		*counter++
		return
	}

	if curr.y-1 >= 0 && trail[curr.y-1][curr.x] == trail[curr.y][curr.x]+1 {
		hikePart2(trail, point{curr.x, curr.y - 1}, counter)
	}
	if curr.x+1 < len(trail[0]) && trail[curr.y][curr.x+1] == trail[curr.y][curr.x]+1 {
		hikePart2(trail, point{curr.x + 1, curr.y}, counter)
	}
	if curr.y+1 < len(trail) && trail[curr.y+1][curr.x] == trail[curr.y][curr.x]+1 {
		hikePart2(trail, point{curr.x, curr.y + 1}, counter)
	}
	if curr.x-1 >= 0 && trail[curr.y][curr.x-1] == trail[curr.y][curr.x]+1 {
		hikePart2(trail, point{curr.x - 1, curr.y}, counter)
	}
}
