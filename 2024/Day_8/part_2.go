package day8_24

import (
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
	data, height, width, game_map := GetData()
	antinodes := make(map[coordinate]bool)
	num_of_antinodes := 0

	for _, coords := range data {
		for i := range coords {
			_, ok := antinodes[coords[i]]
			if !ok {
				antinodes[coords[i]] = true
			}
		}
	}

	for _, coords := range data {
		num_of_antinodes += len(coords)
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				x_diff := coords[j].x - coords[i].x
				y_diff := coords[j].y - coords[i].y

				antinode := coords[i]
				for {
					antinode = coordinate{
						x: antinode.x - x_diff,
						y: antinode.y - y_diff,
					}

					_, ok := antinodes[antinode]
					if !ok {
						if antinode.x >= 0 && antinode.y >= 0 && antinode.x <= width && antinode.y <= height {
							antinodes[antinode] = true
							num_of_antinodes++
							game_map[antinode.y][antinode.x] = '#'
						} else {
							break
						}
					}
				}

				antinode = coords[j]
				for {
					antinode = coordinate{
						x: antinode.x + x_diff,
						y: antinode.y + y_diff,
					}

					_, ok := antinodes[antinode]
					if !ok {
						if antinode.x >= 0 && antinode.y >= 0 && antinode.x <= width && antinode.y <= height {
							antinodes[antinode] = true
							num_of_antinodes++
							game_map[antinode.y][antinode.x] = '#'
						} else {
							break
						}
					}
				}
			}
		}
	}

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + strconv.Itoa(num_of_antinodes))
	shared.ErrCheck(err)
	out.Sync()
}
