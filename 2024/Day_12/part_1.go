package day12_24

import (
	"fmt"
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

type coordinate struct {
	y int
	x int
}

func Part1() {
	garden := GetGarden()
	seen := make(map[coordinate]bool)
	price := 0
	area := 0
	perimeter := 0

	for i, row := range garden {
		for j := range row {
			cord := coordinate{i, j}
			_, ok := seen[cord]
			if !ok {
				Plot(garden, cord, &seen, &area, &perimeter)
				price += area * perimeter
				area, perimeter = 0, 0
			}
		}
	}

	fmt.Println(price)

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(price))
	shared.ErrCheck(err)
	out.Sync()
}

func Plot(garden [][]byte, curr coordinate, seen *map[coordinate]bool, area *int, perimeter *int) {
	temp_seen := *seen
	_, ok := temp_seen[curr]
	if ok {
		return
	}

	temp_seen[curr] = true
	*seen = temp_seen
	*area++
	adjacent := 0

	if curr.y-1 >= 0 && garden[curr.y-1][curr.x] == garden[curr.y][curr.x] {
		adjacent++
	}
	if curr.x+1 < len(garden[0]) && garden[curr.y][curr.x+1] == garden[curr.y][curr.x] {
		adjacent++
	}
	if curr.y+1 < len(garden) && garden[curr.y+1][curr.x] == garden[curr.y][curr.x] {
		adjacent++
	}
	if curr.x-1 >= 0 && garden[curr.y][curr.x-1] == garden[curr.y][curr.x] {
		adjacent++
	}

	*perimeter += 4 - adjacent

	if curr.y-1 >= 0 && garden[curr.y-1][curr.x] == garden[curr.y][curr.x] {
		Plot(garden, coordinate{curr.y - 1, curr.x}, seen, area, perimeter)
	}
	if curr.x+1 < len(garden[0]) && garden[curr.y][curr.x+1] == garden[curr.y][curr.x] {
		Plot(garden, coordinate{curr.y, curr.x + 1}, seen, area, perimeter)
	}
	if curr.y+1 < len(garden) && garden[curr.y+1][curr.x] == garden[curr.y][curr.x] {
		Plot(garden, coordinate{curr.y + 1, curr.x}, seen, area, perimeter)
	}
	if curr.x-1 >= 0 && garden[curr.y][curr.x-1] == garden[curr.y][curr.x] {
		Plot(garden, coordinate{curr.y, curr.x - 1}, seen, area, perimeter)
	}

	return
}

func GetGarden() [][]byte {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	cursor, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	row := []byte{}
	out := [][]byte{}

	for cursor < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		cursor++

		if string(char[0]) == "\n" {
			out = append(out, row)
			row = []byte{}
			continue
		}

		row = append(row, char[0])
	}

	return out
}
