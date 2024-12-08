package day6_24

import (
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

type coordinate struct {
	i int
	j int
}

func Part2() {
	level_map, start := GetMapAndStart()
	path_directions := make(map[coordinate][]string)
	num_of_walls := 0
	path_found := false
	wall_placed := false

	i := start[0]
	j := start[1]

	for {
		//UP
		for level_map[i][j] != '#' {
			k := j
			for k < len(level_map[0]) {
				if level_map[i][k] == '#' {
					cord := coordinate{
						i: i,
						j: k - 1,
					}
					dirs, ok := path_directions[cord]
					if ok {
						for _, dir := range dirs {
							if dir == "right" || dir == "down" {
								num_of_walls++
								wall_placed = true
								break
							} else {
								if IsLoop(level_map, []int{i, k - 1}, "right") {
									num_of_walls++
									wall_placed = true
									break
								}
							}
						}
						if wall_placed {
							break
						}
					}
				}
				k++
			}
			cord := coordinate{
				i: i,
				j: j,
			}
			path_directions[cord] = append(path_directions[cord], "up")
			i--
			if i < 0 {
				path_found = true
				break
			}
		}
		if path_found {
			break
		}
		wall_placed = false
		i++

		//RIGHT
		for level_map[i][j] != '#' {
			k := i
			for k < len(level_map) {
				if level_map[k][j] == '#' {
					cord := coordinate{
						i: k - 1,
						j: j,
					}
					dirs, ok := path_directions[cord]
					if ok {
						for _, dir := range dirs {
							if dir == "down" || dir == "left" {
								num_of_walls++
								wall_placed = true
								break
							} else {
								if IsLoop(level_map, []int{k - 1, j}, "down") {
									num_of_walls++
									wall_placed = true
									break
								}
							}
						}
						if wall_placed {
							break
						}
					}
				}
				k++
			}
			cord := coordinate{
				i: i,
				j: j,
			}
			path_directions[cord] = append(path_directions[cord], "right")
			j++
			if j >= len(level_map[0]) {
				path_found = true
				break
			}
		}
		if path_found {
			break
		}
		wall_placed = false
		j--

		//DOWN
		for level_map[i][j] != '#' {
			k := j
			for k >= 0 {
				if level_map[i][k] == '#' {
					cord := coordinate{
						i: i,
						j: k + 1,
					}
					dirs, ok := path_directions[cord]
					if ok {
						for _, dir := range dirs {
							if dir == "left" || dir == "up" {
								num_of_walls++
								wall_placed = true
								break
							} else {
								if IsLoop(level_map, []int{i, k + 1}, "left") {
									num_of_walls++
									wall_placed = true
									break
								}
							}
						}
						if wall_placed {
							break
						}
					}
				}
				k--
			}
			cord := coordinate{
				i: i,
				j: j,
			}
			path_directions[cord] = append(path_directions[cord], "down")
			i++
			if i >= len(level_map) {
				path_found = true
				break
			}
		}
		if path_found {
			break
		}
		wall_placed = false
		i--

		//LEFT
		for level_map[i][j] != '#' {
			k := i
			for k >= 0 {
				if level_map[k][j] == '#' {
					cord := coordinate{
						i: k + 1,
						j: j,
					}
					dirs, ok := path_directions[cord]
					if ok {
						for _, dir := range dirs {
							if dir == "up" || dir == "right" {
								num_of_walls++
								wall_placed = true
								break
							} else {
								if IsLoop(level_map, []int{k + 1, j}, "up") {
									num_of_walls++
									wall_placed = true
									break
								}
							}
						}
						if wall_placed {
							break
						}
					}
				}
				k--
			}
			cord := coordinate{
				i: i,
				j: j,
			}
			path_directions[cord] = append(path_directions[cord], "left")
			j--
			if j < 0 {
				path_found = true
				break
			}
		}
		if path_found {
			break
		}
		wall_placed = false
		j++
	}

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + strconv.Itoa(num_of_walls))
	shared.ErrCheck(err)
	out.Sync()
}

func IsLoop(level_map [][]byte, start []int, start_dir string) bool {
	i := start[0]
	j := start[1]
	path_found := false
	path := make(map[coordinate][]string)

	for {
		if start_dir == "up" || start_dir == "" {
			for level_map[i][j] != '#' {
				cord := coordinate{i, j}
				dirs, ok := path[cord]
				if ok {
					for _, dir := range dirs {
						if dir == "up" {
							return true
						}
					}
				}
				path[cord] = append(path[cord], "up")
				start_dir = ""
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
		}
		if start_dir == "right" || start_dir == "" {
			for level_map[i][j] != '#' {
				cord := coordinate{i, j}
				dirs, ok := path[cord]
				if ok {
					for _, dir := range dirs {
						if dir == "right" {
							return true
						}
					}
				}
				path[cord] = append(path[cord], "right")
				start_dir = ""
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
		}
		if start_dir == "down" || start_dir == "" {
			for level_map[i][j] != '#' {
				cord := coordinate{i, j}
				dirs, ok := path[cord]
				if ok {
					for _, dir := range dirs {
						if dir == "down" {
							return true
						}
					}
				}
				path[cord] = append(path[cord], "down")
				start_dir = ""
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
		}
		if start_dir == "left" || start_dir == "" {
			for level_map[i][j] != '#' {
				cord := coordinate{i, j}
				dirs, ok := path[cord]
				if ok {
					for _, dir := range dirs {
						if dir == "left" {
							return true
						}
					}
				}
				path[cord] = append(path[cord], "left")
				start_dir = ""
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
	}

	return false
}
