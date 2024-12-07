package day6_24

import (
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
	level_map, start := GetMapAndStart()
	seen_path := [][]bool{}
	num_of_walls := 0
	path_found := false
	placed_wall := false

	for _, arr := range level_map {
		temp_arr := []bool{}
		for range arr {
			temp_arr = append(temp_arr, false)
		}
		seen_path = append(seen_path, temp_arr)
	}

	i := start[0]
	j := start[1]

	for {
		//UP
		for level_map[i][j] != '#' {
			k := j
			for k < len(level_map[0]) {
				if level_map[i][k] == '#' && seen_path[i][k-1] {
					l := i
					for l < len(level_map) {
						if level_map[l][k-1] == '#' && seen_path[l-1][k-1] {
							num_of_walls++
							placed_wall = true
							break
						}
						l++
					}
					if placed_wall {
						placed_wall = false
						break
					}
				}
				k++
			}
			seen_path[i][j] = true
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

		//RIGHT
		for level_map[i][j] != '#' {
			k := i
			for k < len(level_map) {
				if level_map[k][j] == '#' && seen_path[k-1][j] {
					l := j
					for l >= 0 {
						if level_map[k-1][l] == '#' && seen_path[k-1][l+1] {
							num_of_walls++
							placed_wall = true
							break
						}
						l--
					}
					if placed_wall {
						placed_wall = false
						break
					}
				}
				k++
			}
			seen_path[i][j] = true
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

		//DOWN
		for level_map[i][j] != '#' {
			k := j
			for k >= 0 {
				if level_map[i][k] == '#' && seen_path[i][k+1] {
					l := i
					for l >= 0 {
						if level_map[l][k+1] == '#' && seen_path[l+1][k+1] {
							num_of_walls++
							placed_wall = true
							break
						}
						l--
					}
					if placed_wall {
						placed_wall = false
						break
					}
				}
				k--
			}
			seen_path[i][j] = true
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

		//LEFT
		for level_map[i][j] != '#' {
			k := i
			for k >= 0 {
				if level_map[k][j] == '#' && seen_path[k+1][j] {
					l := j
					for l < len(level_map[0]) {
						if level_map[k+1][l] == '#' && seen_path[k+1][l-1] {
							num_of_walls++
							placed_wall = true
							break
						}
						l++
					}
					if placed_wall {
						placed_wall = false
						break
					}
				}
				k--
			}
			seen_path[i][j] = true
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

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + strconv.Itoa(num_of_walls))
	shared.ErrCheck(err)
	out.Sync()
}
