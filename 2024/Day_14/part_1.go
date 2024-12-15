package day14_24

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	shared "github.com/VladanT3/Advent_of_Code"
)

type coordinate struct {
	x int
	y int
}

type velocity struct {
	vel_x int
	vel_y int
}

func Part1() {
	robots := GetRobots()
	new_robots := make(map[coordinate][]velocity)

	for range 100 {
		for pos, vels := range robots {
			for _, vel := range vels {
				new_x := pos.x + vel.vel_x
				new_y := pos.y + vel.vel_y
				if new_x > 100 {
					new_x = new_x - 100
				} else if new_x < 0 {
					new_x = new_x + 100
				}
				if new_y > 102 {
					new_y = new_y - 102
				} else if new_y < 0 {
					new_y = new_y + 102
				}
				new_pos := coordinate{new_x, new_y}
				if new_robots[new_pos] == nil {
					new_robots[new_pos] = []velocity{}
				}
				new_robots[new_pos] = append(new_robots[new_pos], vel)
			}
		}
		robots = new_robots
		new_robots = make(map[coordinate][]velocity)
	}

	quadrant_1 := big.NewInt(0)
	quadrant_2 := big.NewInt(0)
	quadrant_3 := big.NewInt(0)
	quadrant_4 := big.NewInt(0)
	for pos, vel := range robots {
		if pos.x < 50 && pos.y < 51 {
			quadrant_1 = new(big.Int).Add(quadrant_1, big.NewInt(int64(len(vel))))
		}
		if pos.x > 50 && pos.y < 51 {
			quadrant_2 = new(big.Int).Add(quadrant_2, big.NewInt(int64(len(vel))))
		}
		if pos.x < 50 && pos.y > 51 {
			quadrant_3 = new(big.Int).Add(quadrant_3, big.NewInt(int64(len(vel))))
		}
		if pos.x > 50 && pos.y > 51 {
			quadrant_4 = new(big.Int).Add(quadrant_4, big.NewInt(int64(len(vel))))
		}
	}

	safety_factor := new(big.Int).Mul(
		new(big.Int).Mul(quadrant_1, quadrant_2),
		new(big.Int).Mul(quadrant_3, quadrant_4))
	fmt.Println(safety_factor)

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + safety_factor.String())
	shared.ErrCheck(err)
	out.Sync()
}

func GetRobots() map[coordinate][]velocity {
	f, err := os.ReadFile("input.txt")
	shared.ErrCheck(err)

	out := make(map[coordinate][]velocity)

	rows := strings.Split(string(f), "\n")
	for i := 0; i < len(rows)-1; i++ {
		row := strings.Split(rows[i], " ")
		pos := strings.Split(row[0], "=")
		vel := strings.Split(row[1], "=")
		pos = strings.Split(pos[1], ",")
		vel = strings.Split(vel[1], ",")

		x, err := strconv.Atoi(pos[0])
		shared.ErrCheck(err)
		y, err := strconv.Atoi(pos[1])
		shared.ErrCheck(err)
		cord := coordinate{x, y}

		vel_x, err := strconv.Atoi(vel[0])
		shared.ErrCheck(err)
		vel_y, err := strconv.Atoi(vel[1])
		shared.ErrCheck(err)
		vels := velocity{vel_x, vel_y}

		if out[cord] == nil {
			out[cord] = []velocity{}
		}
		out[cord] = append(out[cord], vels)
	}

	return out
}
