package day8_24

import (
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

type coordinate struct {
	x int
	y int
}

func Part1() {
	data, height, width, game_map := GetData()
	antinodes := make(map[coordinate]bool)
	num_of_antinodes := 0

	for _, coords := range data {
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				x_diff := coords[j].x - coords[i].x
				y_diff := coords[j].y - coords[i].y

				antinode := coordinate{
					x: coords[i].x - x_diff,
					y: coords[i].y - y_diff,
				}

				_, ok := antinodes[antinode]
				if !ok {
					if antinode.x >= 0 && antinode.y >= 0 && antinode.x <= width && antinode.y <= height {
						antinodes[antinode] = true
						num_of_antinodes++
						game_map[antinode.y][antinode.x] = '#'
					}
				}

				antinode = coordinate{
					x: coords[j].x + x_diff,
					y: coords[j].y + y_diff,
				}

				_, ok = antinodes[antinode]
				if !ok {
					if antinode.x >= 0 && antinode.y >= 0 && antinode.x <= width && antinode.y <= height {
						antinodes[antinode] = true
						num_of_antinodes++
						game_map[antinode.y][antinode.x] = '#'
					}
				}
			}
		}
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(num_of_antinodes))
	shared.ErrCheck(err)
	out.Sync()
}

func GetData() (map[byte][]coordinate, int, int, [][]byte) {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	cursor, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	i := 0
	j := 0
	out := make(map[byte][]coordinate)
	height := 0
	width := 0
	game_map := [][]byte{}
	row := []byte{}

	for cursor < eof {
		_, err := f.Read(char)
		shared.ErrCheck(err)
		cursor++

		if string(char[0]) == "\n" {
			width = j - 1
			i++
			j = 0
			game_map = append(game_map, row)
			row = []byte{}
			continue
		}

		row = append(row, char[0])

		if char[0] != '.' {
			if out[char[0]] == nil {
				out[char[0]] = []coordinate{}
			}
			out[char[0]] = append(out[char[0]], coordinate{j, i})
		}
		j++
	}

	height = i - 1

	return out, height, width, game_map
}
