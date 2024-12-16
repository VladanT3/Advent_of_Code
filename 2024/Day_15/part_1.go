package day15_24

import (
	"fmt"
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

type position struct {
	y int
	x int
}

func Part1() {
	warehouse_map, start := GetMap()
	directions, err := os.ReadFile("input.txt")
	shared.ErrCheck(err)
	curr := start
	run_into_wall := false

	for i := range directions {
		if directions[i] == '^' {
			if warehouse_map[curr.y-1][curr.x] == '.' {
				warehouse_map[curr.y-1][curr.x] = '@'
				warehouse_map[curr.y][curr.x] = '.'
				curr.y--
			} else if warehouse_map[curr.y-1][curr.x] == '#' {
				continue
			} else if warehouse_map[curr.y-1][curr.x] == 'O' {
				for j := curr.y - 1; j >= 0; j-- {
					if warehouse_map[j][curr.x] == '#' {
						run_into_wall = true
						break
					} else if warehouse_map[j][curr.x] == '.' {
						warehouse_map[j][curr.x] = 'O'
						break
					}
				}
				if run_into_wall {
					run_into_wall = false
					continue
				} else {
					warehouse_map[curr.y-1][curr.x] = '@'
					warehouse_map[curr.y][curr.x] = '.'
					curr.y--
				}
			}
		} else if directions[i] == '>' {
			if warehouse_map[curr.y][curr.x+1] == '.' {
				warehouse_map[curr.y][curr.x+1] = '@'
				warehouse_map[curr.y][curr.x] = '.'
				curr.x++
			} else if warehouse_map[curr.y][curr.x+1] == '#' {
				continue
			} else if warehouse_map[curr.y][curr.x+1] == 'O' {
				for j := curr.x + 1; j < len(warehouse_map[0]); j++ {
					if warehouse_map[curr.y][j] == '#' {
						run_into_wall = true
						break
					} else if warehouse_map[curr.y][j] == '.' {
						warehouse_map[curr.y][j] = 'O'
						break
					}
				}
				if run_into_wall {
					run_into_wall = false
					continue
				} else {
					warehouse_map[curr.y][curr.x+1] = '@'
					warehouse_map[curr.y][curr.x] = '.'
					curr.x++
				}
			}
		} else if directions[i] == 'v' {
			if warehouse_map[curr.y+1][curr.x] == '.' {
				warehouse_map[curr.y+1][curr.x] = '@'
				warehouse_map[curr.y][curr.x] = '.'
				curr.y++
			} else if warehouse_map[curr.y+1][curr.x] == '#' {
				continue
			} else if warehouse_map[curr.y+1][curr.x] == 'O' {
				for j := curr.y + 1; j < len(warehouse_map); j++ {
					if warehouse_map[j][curr.x] == '#' {
						run_into_wall = true
						break
					} else if warehouse_map[j][curr.x] == '.' {
						warehouse_map[j][curr.x] = 'O'
						break
					}
				}
				if run_into_wall {
					run_into_wall = false
					continue
				} else {
					warehouse_map[curr.y+1][curr.x] = '@'
					warehouse_map[curr.y][curr.x] = '.'
					curr.y++
				}
			}
		} else if directions[i] == '<' {
			if warehouse_map[curr.y][curr.x-1] == '.' {
				warehouse_map[curr.y][curr.x-1] = '@'
				warehouse_map[curr.y][curr.x] = '.'
				curr.x--
			} else if warehouse_map[curr.y][curr.x-1] == '#' {
				continue
			} else if warehouse_map[curr.y][curr.x-1] == 'O' {
				for j := curr.x - 1; j >= 0; j-- {
					if warehouse_map[curr.y][j] == '#' {
						run_into_wall = true
						break
					} else if warehouse_map[curr.y][j] == '.' {
						warehouse_map[curr.y][j] = 'O'
						break
					}
				}
				if run_into_wall {
					run_into_wall = false
					continue
				} else {
					warehouse_map[curr.y][curr.x-1] = '@'
					warehouse_map[curr.y][curr.x] = '.'
					curr.x--
				}
			}
		}
	}

	sum := 0
	for i, row := range warehouse_map {
		for j := range row {
			if row[j] == 'O' {
				sum += 100*i + j
			}
		}
	}

	fmt.Println(sum)

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(sum))
	shared.ErrCheck(err)
	out.Sync()
}

func GetMap() ([][]byte, position) {
	f, err := os.Open("map.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	cursor, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	out := [][]byte{}
	row := []byte{}
	i, j := 0, 0
	found_start := false

	for cursor < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		cursor++

		if char[0] == '@' {
			found_start = true
		}

		if string(char[0]) == "\n" {
			out = append(out, row)
			row = []byte{}
			if !found_start {
				i++
				j = 0
			}
			continue
		}

		row = append(row, char[0])
		if !found_start {
			j++
		}
	}

	return out, position{i, j}
}
