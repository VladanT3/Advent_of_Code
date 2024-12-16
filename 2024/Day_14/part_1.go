package day14_24

import (
	"fmt"
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
	x int
	y int
}

type robot struct {
	pos coordinate
	vel velocity
}

const height = 103
const width = 101

func Part1() {
	robots := GetRobots()

	for range 100 {
		for i := range robots {
			robots[i].pos.x += robots[i].vel.x
			robots[i].pos.y += robots[i].vel.y

			if robots[i].pos.x >= width {
				robots[i].pos.x -= width
			} else if robots[i].pos.x < 0 {
				robots[i].pos.x += width
			}
			if robots[i].pos.y >= height {
				robots[i].pos.y -= height
			} else if robots[i].pos.y < 0 {
				robots[i].pos.y += height
			}
		}
	}

	quadrant_1 := 0
	quadrant_2 := 0
	quadrant_3 := 0
	quadrant_4 := 0
	for _, robot := range robots {
		if robot.pos.x < width/2 && robot.pos.y < height/2 {
			quadrant_1++
		}
		if robot.pos.x > width/2 && robot.pos.y < height/2 {
			quadrant_2++
		}
		if robot.pos.x < width/2 && robot.pos.y > height/2 {
			quadrant_3++
		}
		if robot.pos.x > width/2 && robot.pos.y > height/2 {
			quadrant_4++
		}
	}

	safety_factor := quadrant_1 * quadrant_2 * quadrant_3 * quadrant_4
	fmt.Println(safety_factor)

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(safety_factor))
	shared.ErrCheck(err)
	out.Sync()
}

func GetRobots() []robot {
	f, err := os.ReadFile("input.txt")
	shared.ErrCheck(err)

	out := []robot{}

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

		out = append(out, robot{cord, vels})
	}

	return out
}
